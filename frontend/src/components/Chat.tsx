import React, { useEffect, useState } from "react";
import { Message } from "../types";
import ChatBubble from "./ChatBubble";
import { ethers } from "ethers";

interface Props {
  account?: string;
  weddingContract: ethers.Contract | undefined;
}

const Chat = ({ account, weddingContract }: Props) => {
  const [textareaContent, setTextareaContent] = useState("");
  const [txnStatus, setTxnStatus] = useState<string | null>(null);
  const [messages, setMessages] = useState<Message[]>();

  const getMessages = async () => {
    if (!weddingContract || account) return;

    const messages = await weddingContract.getMessages();

    setMessages(() => {
      return messages.map((w: any) => ({
        address: w.sender,
        date: w.timestamp._hex,
        content: w.content,
      }));
    });
  };

  const setupMessageListener = (): ethers.Contract | void => {
    if (!weddingContract) return;

    const msgListener = weddingContract.on(
      "NewMessage",
      (address, timestamp, content, _style) => {
        setMessages((prev) => {
          const newMessage = {
            address,
            date: timestamp._hex,
            content,
          };
          return prev ? [...prev, newMessage] : [newMessage];
        });
      }
    );

    return msgListener;
  };

  const sendMessage = async () => {
    if (!weddingContract) return;
    try {
      setTxnStatus("WAIT");
      const messageTxn = await weddingContract.sendMessage(textareaContent);
      setTxnStatus("SENDING");
      await messageTxn.wait();
    } catch (e) {
      console.warn("Transaction failed with error", e);
    } finally {
      setTextareaContent("");
      setTxnStatus(null);
    }
  };

  useEffect(() => {
    if (!weddingContract || messages) return;
    getMessages();
    setupMessageListener();
  }, [weddingContract]);


  // console.log("messages", messages);

  return (
    <div className="chat">
      <h2>Wish them the best</h2>
      <hr></hr>
      <div className="chat__messages">
        {!weddingContract && (
          <p className="state-message">
            Connect to the chat in order to see the messages!
          </p>
        )}
        {account && messages && messages.length === 0 && (
          <p className="state-message">There is no message to display</p>
        )}
        {messages &&
          messages.length > 0 &&
          messages.map((m, i) => (
            <ChatBubble
              key={i}
              ownMessage={m.address === account}
              address={m.address}
              message={m.content}
            />
          ))}
      </div>
      <div className="chat__actions-wrapper">
        {!account && (
          <p className="state-message">Connect your wallet to send a wish!</p>
        )}
        <div className="chat__input">
          <textarea
            disabled={!!txnStatus || !account}
            value={textareaContent}
            onChange={(e) => {
              setTextareaContent(e.target.value);
            }}
          ></textarea>
          <button onClick={sendMessage} disabled={!!txnStatus || !account}>
            {txnStatus || "Send"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default Chat;
