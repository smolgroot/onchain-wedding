// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;


contract Wedding {


    // // WEDDING LOGIC
    // // ------------------------------------------------------

    struct Fiance  {
        address sender;
        string description;
        uint timestamp;
    }

    // Manage those vars in a constructor
    Fiance fiance1;
    Fiance fiance2;
    
    Fiance[] signers;

    event NewSignature(address indexed from, uint timestamp);

    function sayYes() public {
        // @TODO --> check if addy is one from fiance1/2 
        signers.push(fiance1);
        emit NewSignature(msg.sender, block.timestamp);
    }

    function updateFianceDescription(Fiance memory fiance) public {
        // TODO
    }


    // DIVORCE LOGIC
    // ------------------------------------------------------
    // TODO

    // GIFT LOGIC
    // ------------------------------------------------------
    // TODO: ether, ERC-20 tokens, ERC-721 NFT, etc.

    // ON-CHAIN Chat
    // ------------------------------------------------------
    // @TODO --> Use XMTP or Push

    event NewMessage(address indexed from, uint timestamp, string message);

    struct Message  {
        address sender;
        string content;
        uint timestamp;
    }

    Message[] messages;

    function sendMessage(string calldata _content) public {
        messages.push(Message(msg.sender, _content, block.timestamp));
        emit NewMessage(msg.sender, block.timestamp, _content);
    }

    function getMessages() view public returns (Message[] memory) {
        return messages;
    }
}