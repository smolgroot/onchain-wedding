import React, { useState } from "react";
import "./App.css";
import Chat from "./components/Chat";
import Wedding from "./contract/contract.json";
import useWeddingContract from "./useWeddingContract";
import Fireworks from './components/Fireworks';
import Attestation from './components/Attestation';
import Ceremony from './components/Ceremony';
import Header from "./components/Header";
import Footer from "./components/Footer";
import Grid from '@mui/material/Grid';
import Marriage from "./components/Marriage";


function App() {
  const contractAddress = "0xf2ed405bf4164ae48ad9fccd90c4ff1622e44565";
  const network = "sepolia";
  const appName = "Smart Wedding";
  const [account, setAccount] = useState<string>();
  const weddingContract = useWeddingContract(
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
          <Chat account={account} weddingContract={weddingContract} />
        </Grid>
        <Grid item xs={7} md={7} className="ceremony">
          <div>.</div>
          <Attestation />
          <Marriage />
          <Ceremony weddingContract={weddingContract} />
          {/* <Fireworks/> */}
        </Grid>
      </Grid>
      <Footer contractAddress={contractAddress}/>
    </div>
  );
}

export default App;
