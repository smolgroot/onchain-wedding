import React, { useState } from "react";
import "./App.css";
import Chat from "./components/Chat";
import Wedding from "./contract/contract.json";
import useWeddingContract from "./useWeddingContract";
import Fireworks from './components/Fireworks';
import Ceremony from './components/Ceremony';
import Header from "./components/Header";
import Footer from "./components/Footer";
import Grid from '@mui/material/Grid';


function App() {
  const contractAddress = "0x4fcd2d1d4175262f110b0c7a055ea25c2aed656c";
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
          <Chat account={account} weddingContract={weddingContract} />
        </Grid>
        <Grid item xs={7} md={7} className="ceremony">
          <Ceremony account={account} weddingContract={weddingContract} />
          {/* <Fireworks/> */}
        </Grid>
      </Grid>
      <Footer contractAddress={contractAddress}/>
    </div>
  );
}

export default App;
