// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;


contract Wedding {

    address public lastSigner;

    // WEDDING LOGIC
    // ------------------------------------------------------

    struct Fiance  {
        address walletAddr;
        string firstname;
        string description;
        string ipfsPhotoCid;
        bool alreadySaidYes;
    }

    // Hardcoded fiances addresses here
    Fiance fiance1 = Fiance(
        0xE42749aA40040929cd8db5ac5200df12f9B52540, 
        "Rajeev", 
        "Hello there, i'm about the get married", 
        "QmSQTz4CxHmrNapTjYZcYdT46qJRjJN29XFj2PtrvFp6Lv",
        false
    );

    Fiance fiance2 = Fiance(
        0x23bC9eB2343125b91a7d5c59038e40E0AE451551, 
        "Maria", 
        "This is the most exciting day in my life!", 
        "QmX4aCHjcuqpDysJk6WDLj69sYTtMGhyTypLu2uwLGKNYg",
        false
    );
    
    Fiance[] signers;

    event NewSignature(address indexed from, uint timestamp);
    event NotAllowed(address indexed from, uint timestamp);
    event MarriageDone(uint timestamp);

    function sayYes() public {

        // Check if sender is one of the 2 fiances
        if (msg.sender == fiance1.walletAddr) {
            fiance1.alreadySaidYes = true;
            signers.push(fiance1);
            emit NewSignature(fiance1.walletAddr, block.timestamp);
        } else if (msg.sender == fiance2.walletAddr)  {
            fiance2.alreadySaidYes = true;
            signers.push(fiance2);
            emit NewSignature(fiance2.walletAddr, block.timestamp);
        } else {
            // if not one of the 2 fiances, emit an NotAllowed event
            emit NotAllowed(msg.sender, block.timestamp);
        }

        // Check if both peers accept the wedding contract
        // If both said "Yes", then emit an event
        if (getSignersCount() == 2) {
            lastSigner = signers[1].walletAddr;
            emit MarriageDone(block.timestamp);
        }   

    }

    function getSignersCount() public view returns(uint count) {
        return signers.length;
    }

    function getLastSigner() public view returns(address) {
        return lastSigner;
    }


    function getFiance1() public view returns(address, string memory) {
        return (fiance1.walletAddr, fiance1.firstname);
    }

    function getFiance2() public view returns(address, string memory) {
        return (fiance2.walletAddr, fiance2.firstname);
    }

    function updateFianceDescription(Fiance memory fiance) public {
        // TODO
    }

    function updateFiancePhoto(Fiance memory fiance) public {
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