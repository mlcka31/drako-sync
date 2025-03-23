// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AlephGameState {
    // Game state definitions.
    enum GameState {
        UserAction,
        AgentAction,
        Complete,
        NotStarted
    }

    GameState public gameState;

    // Prize pool holds accumulated funds.
    uint256 public prizePool;
    // Base fee for sending a message.
    uint256 public messagePrice;
    // Coefficient used to increase the message fee (fixed-point, scaled by 1e18).
    uint256 public coeffIncrease;
    // Scaling factor for fixed-point arithmetic.
    uint256 public constant PRECISION = 1000000000000000000;

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
    constructor(uint256 _messagePrice, uint256 _coeffIncrease) {
        adminAddress = msg.sender;
        messagePrice = _messagePrice;
        coeffIncrease = _coeffIncrease;
        gameState = GameState.NotStarted;
    }

    // Getter and Setter for gameState
    function getGameState() public view returns (GameState) {
        return gameState;
    }

    function setGameState(GameState _gameState) public onlyAdmin {
        gameState = _gameState;
    }

    // Getter and Setter for prizePool
    function getPrizePool() public view returns (uint256) {
        return prizePool;
    }

    function setPrizePool(uint256 _prizePool) public onlyAdmin {
        prizePool = _prizePool;
    }

    // Getter and Setter for messagePrice
    function getmessagePrice() public view returns (uint256) {
        return messagePrice;
    }

    function setmessagePrice(uint256 _messagePrice) public onlyAdmin {
        messagePrice = _messagePrice;
    }

    // Getter and Setter for coeffIncrease
    function getCoeffIncrease() public view returns (uint256) {
        return coeffIncrease;
    }

    function setCoeffIncrease(uint256 _coeffIncrease) public onlyAdmin {
        coeffIncrease = _coeffIncrease;
    }

    // Getter for PRECISION (constant, no setter needed)
    function getPrecision() public pure returns (uint256) {
        return PRECISION;
    }

    // Setter for adminAddress
    function setAdminAddress(address _adminAddress) public onlyAdmin {
        adminAddress = _adminAddress;
    }

    // Setter for aiAgentAddress
    function setAIAgentAddress(address _aiAgentAddress) public onlyAdmin {
        aiAgentAddress = _aiAgentAddress;
        emit AIAgentSetAddress(_aiAgentAddress);
    }

    // Setter for AI agent address; callable only by admin.
    function initGame(address _aiAgentAddress) external onlyAdmin {
        require(aiAgentAddress == address(0), "Game already initialized");

        aiAgentAddress = _aiAgentAddress;
        emit AIAgentSetAddress(_aiAgentAddress);
        gameState = GameState.UserAction;
    }

    // Users call this function to send a message and pay the fee.
    function sendMessage(string memory _content) public payable {
        require(gameState == GameState.UserAction, "Game is not active or it is not your turn");
        require(msg.value == messagePrice, "Wrong message price");

        // Record the user message.
        messages.push(Message(messagePrice, msg.sender, _content, block.timestamp));
        // Increase the prize pool.
        prizePool += msg.value;
        // Increase the message fee using fixed-point multiplication:
        messagePrice = (messagePrice * coeffIncrease) / PRECISION;

        emit MessageSent(msg.sender, _content, block.timestamp);

        gameState = GameState.AgentAction;
    }

    // AI agent calls this function to reply.
    function reply(string memory _content) public {
        require(msg.sender == aiAgentAddress, "Only AI agent can reply");
        require(gameState == GameState.AgentAction, "Game is not in going state");

        // Record the agent's reply.
        messages.push(Message(0, aiAgentAddress, _content, block.timestamp));
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

    function abort() public onlyAdmin {
        uint256 totalRefunded = 0;

        // Loop over all messages and refund players their contributions.
        for (uint256 i = 0; i < messages.length; i++) {
            // Exclude messages from the AI agent and the admin.
            if (messages[i].sender != aiAgentAddress && messages[i].sender != adminAddress) {
                uint256 refund = messages[i].messagePrice;
                totalRefunded += refund;
                // Refund the sender. Using call to transfer funds.
                (bool success,) = messages[i].sender.call{value: refund}("");
                require(success, "Refund transfer failed");
            }
        }

        // Calculate admin's share: the remaining funds after refunds.
        uint256 adminShare = prizePool - totalRefunded;
        // Set prizePool to zero before transferring to prevent reentrancy issues.
        prizePool = 0;
        (bool sent,) = adminAddress.call{value: adminShare}("");
        require(sent, "Admin transfer failed");

        // Update the game state to complete.
        gameState = GameState.Complete;
    }

    // Function to retrieve all messages.
    function getMessages() public view returns (Message[] memory) {
        return messages;
    }
}
