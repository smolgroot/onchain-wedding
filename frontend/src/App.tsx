import React, { useState } from "react";
import "./App.css";
import Chat from "./components/Chat";
import Wedding from "./contract/contract.json";
import useChatContract from "./useChatContract";
import Fireworks from './components/Fireworks';
import Attestation from './components/Attestation';
import Ceremony from './components/Ceremony';
import Header from "./components/Header";
import Footer from "./components/Footer";
import Grid from '@mui/material/Grid';
import Marriage from "./components/Marriage";


function App() {
  const contractAddress = "0x8dd43110d79d51d29155a94e2f73754f86f13981";
  const network = "sepolia";
  const appName = "Smart Wedding";
  const [account, setAccount] = useState<string>();
  const chatContract = useChatContract(
    contractAddress,
    Wedding.abi,
    account
  );

  return (

    <div className="App">
      <Grid container>
        <Header setAccount={setAccount} account={account} contractAddress={contractAddress} appName={appName} network={network}/>
      </Grid>
      <Grid container>
        <Grid item xs={5} md={5} >
          <div>XMTP live chat here</div>
          <Chat account={account} chatContract={chatContract} />
        </Grid>
        <Grid item xs={7} md={7} className="ceremony">
          <div>.</div>
          <Attestation />
          <Marriage />
          <Ceremony account={account} chatContract={chatContract} />
          {/* <Fireworks/> */}
        </Grid>
      </Grid>
      <Footer setAccount={setAccount} account={account} contractAddress={contractAddress} appName={appName}/>
    </div>
  );
}

export default App;
