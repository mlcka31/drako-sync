// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AlephGameState {
    // Game state definitions.
    enum GameState {UserAction, AgentAction, Complete, NotStarted}
    GameState public gameState;

    // Prize pool holds accumulated funds.
    uint256 public prizePool;
    // Base fee for sending a message.
    uint256 public messagePrize;
    // Coefficient used to increase the message fee (fixed-point, scaled by 1e18).
    uint256 public coeffIncrease;
    // Scaling factor for fixed-point arithmetic.
    uint256 public constant PRECISION = 1e18;

    // Structure to store messages.
    struct Message {
        uint256 messagePrice;
        address sender;
        string content;
        uint256 timestamp;
    }

    Message[] public messages;

    // Addresses for admin and AI agent.
    address public adminAddress;
    address public aiAgentAddress;

    // Events for off-chain tracking.
    event MessageSent(address indexed sender, string content, uint256 timestamp);
    event AgentReplied(string content, uint256 timestamp);
    event PrizePayout(address winner, uint256 amount);
    event AIAgentSetAddress(address aiAgent);

    // Modifier to restrict function access to admin.
    modifier onlyAdmin() {
        require(msg.sender == adminAddress, "Only admin can call this function");
        _;
    }

    // Constructor: sets initial fee, coefficient, and game state.
    // Note: coeffIncrease should be passed as a fixed-point number.
    // For example, for a 10% increase, pass 1100000000000000000 (which equals 1.1 * PRECISION)
    constructor(uint256 _messagePrize, uint256 _coeffIncrease) {
        adminAddress = msg.sender;
        messagePrize = _messagePrize;
        coeffIncrease = _coeffIncrease;
        gameState = GameState.NotStarted;
    }

    // Setter for AI agent address; callable only by admin.
    function initGame(address _aiAgentAddress) external onlyAdmin {
        aiAgentAddress = _aiAgentAddress;
        emit AIAgentSetAddress(_aiAgentAddress);
        gameState = GameState.UserAction;
    }

    // Users call this function to send a message and pay the fee.
    function sendMessage(string memory _content) public payable {
        require(gameState == GameState.AgentAction, "Game is not active");
        require(msg.value >= messagePrize, "Insufficient funds for message");

        // Record the user message.
        messages.push(Message(
            messagePrize,
            msg.sender,
            _content,
            block.timestamp));
        // Increase the prize pool.
        prizePool += msg.value;
        // Increase the message fee using fixed-point multiplication:
        // newMessagePrize = (oldMessagePrize * coeffIncrease) / PRECISION.
        messagePrize = (messagePrize * coeffIncrease) / PRECISION;

        emit MessageSent(msg.sender, _content, block.timestamp);
    }

    // AI agent calls this function to reply.
    function reply(string memory _content) public {
        require(msg.sender == aiAgentAddress, "Only AI agent can reply");
        require(gameState == GameState.AgentAction, "Game is not in going state");

        // Record the agent's reply.
        messages.push(Message(
            0,
            aiAgentAddress,
            _content,
            block.timestamp
        ));
        // Transition the game state to Pending.
        gameState = GameState.UserAction;

        emit AgentReplied(_content, block.timestamp);
    }

    // Admin calls this function to payout the prize.
    // For demonstration, the winner is chosen as the last user (non-AI) who sent a message.
    function payoutPrize() public onlyAdmin {
        require(gameState == GameState.UserAction, "Game is not pending payout");
        require(messages.length > 0, "No messages available");

        address winner;
        // If the last message was from the AI agent, pick the message before it.
        if (messages[messages.length - 1].sender == aiAgentAddress && messages.length > 1) {
            winner = messages[messages.length - 2].sender;
        } else {
            winner = messages[messages.length - 1].sender;
        }

        uint256 amount = prizePool;
        prizePool = 0;
        // Transition game state to End.
        gameState = GameState.Complete;

        // Transfer the entire prize pool to the winner.
        (bool sent,) = winner.call{value: amount}("");
        require(sent, "Transfer failed");

        emit PrizePayout(winner, amount);
    }

    // Function to retrieve all messages.
    function getMessages() public view returns (Message[] memory) {
        return messages;
    }
}