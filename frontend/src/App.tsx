import React, { useState } from "react";
import "./App.css";
import Chat from "./components/Chat";
import Sidebar from "./components/Sidebar";
import Wedding from "./contract/contract.json";
import useChatContract from "./useChatContract";
import Fireworks from './components/Fireworks';
import Attestation from './components/Attestation';


function App() {
  const contractAddress = "0xe68C1C1B6316507410FA5fC4E0Ab0400eECE30a17a8bc70";
  const [account, setAccount] = useState<string>();
  const chatContract = useChatContract(
    contractAddress,
    Wedding.abi,
    account
  );

  return (
    <div className="App">
      <Attestation/>
      <Fireworks/>
      <Sidebar setAccount={setAccount} account={account} contractAddress={contractAddress}/>
      <Chat account={account} chatContract={chatContract} />
    </div>
  );
}

export default App;
