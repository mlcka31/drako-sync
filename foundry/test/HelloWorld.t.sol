// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {HelloWorld} from "../src/HelloWorld.sol";

contract HelloWorldTest is Test {
    HelloWorld private helloWorld;

    function setUp() public {
        helloWorld = new HelloWorld();
    }

    function testGreeting() public view {
        string memory expectedGreeting = "Hello, World!";
        assertEq(helloWorld.greeting(), expectedGreeting);
    }

    function testSetGreeting() public {
        string memory expected = "Hello, Bob!";
        helloWorld.setGreeting(expected);
        assertEq(helloWorld.greeting(), expected);
        assertEq(helloWorld.iteration(), 1);
    }
}
