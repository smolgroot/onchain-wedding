import React, { useState } from "react";
import "./App.css";
import Chat from "./components/Chat";
import Sidebar from "./components/Sidebar";
import Wedding from "./contract/contract.json";
import useChatContract from "./useChatContract";

function App() {
  const contractAddress = "0x91cae95fc8c7e5e332982af2eb0bc32327a8bc70";
  const [account, setAccount] = useState<string>();
  const chatContract = useChatContract(
    contractAddress,
    Wedding.abi,
    account
  );

  return (
    <div className="App">
      <Sidebar setAccount={setAccount} account={account} contractAddress={contractAddress}/>
      <Chat account={account} chatContract={chatContract} />
    </div>
  );
}

export default App;
