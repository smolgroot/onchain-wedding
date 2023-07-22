import { ethers } from 'ethers';
import { WeddingDetails } from "../types";
import React, { useEffect, useState } from 'react';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import { Avatar } from '@mui/material';
import { Grid } from '@mui/material';
import { Button } from '@mui/material';



interface Props {
  weddingContract: ethers.Contract | undefined;
}


const Ceremony = ({ weddingContract }: Props) => {
  // const [wd, getFianceDetails] = useState<WeddingDetails[]>();
  const [f1a, setFiance1Addr] = useState<string>();

  const getInformation = async () => {
    if (!weddingContract) return;

    // Fetch fiance details:
    const getFianceDetails = async () => {
      const [fiance1WalletAddr, fiance1Name] = await weddingContract.getFiance1();
      const [fiance2WalletAddr, fiance2Name] = await weddingContract.getFiance2();
      setFiance1Addr(Address1);

      return {
        Address1: fiance1WalletAddr,
        Name1: fiance1Name,
        Address2: fiance2WalletAddr,
        Name2: fiance2Name
      };
    }

    const { Address1, Name1, Address2, Name2 } = await getFianceDetails();
    console.log(Address1, Name1, Address2, Name2);
    setFiance1Addr(Address1);
  };

  // useEffect(() => {
  //   if (!weddingContract ) return;
  // }, [weddingContract]);

  return (
    <div>
      <Grid container id="greetings">
          <h1>Welcome everyone!</h1>
      </Grid>
      <Grid container>
        <Grid item xs={6} md={6} >
        

          <div className="weddingCard">
            <h2>Partner 1</h2>
            <p>Some description about me</p>
            <Avatar
              className='weddingAvatar'
              alt="fiance"
              src="https://bafybeiebtleqr4mtvuq273umq3v637mksfloul2zsfdlvxuwp6hmvubrie.ipfs.dweb.link/"
              sx={{ width: 110, height: 110 }}
            />
            <hr></hr>
            <Button className='weddingButton' variant="contained"  color="success" endIcon={<CheckCircleIcon />}>
              Say Yes
            </Button>
          </div>


        </Grid>
        <Grid item xs={6} md={6}>

        <div className="weddingCard">
          <h2>Partner 1</h2>
          <p>Some description about me</p>
          <Avatar
            className='weddingAvatar'
            alt="fiance"
            src="https://bafybeib4ndy3fuysk7fcjo6yxycdans4zbyzpnmajwfo45n7ufyd2aywj4.ipfs.dweb.link/"
            sx={{ width: 110, height: 110 }}
          />
          <hr></hr>
          <Button className='weddingButton' variant="contained"  color="success" endIcon={<CheckCircleIcon />}>
            Say Yes
          </Button>
        </div>


        </Grid>
      </Grid>
      
      {f1a}
      {/* Today, {name1} and {name2} are getting married! */}
    </div>
  );
};

export default Ceremony;
