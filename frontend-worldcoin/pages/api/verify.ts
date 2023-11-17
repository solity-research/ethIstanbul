import { NextApiRequest, NextApiResponse } from "next";

export type VerifyReply = {
    code: string;
    detail?: string;
  };
  
  export default function handler(req: NextApiRequest, res: NextApiResponse<VerifyReply>) {
    const reqBody = {
      merkle_root: req.body.merkle_root,
      nullifier_hash: req.body.nullifier_hash,
      proof: req.body.proof,
      credential_type: req.body.credential_type,
      action: "login_eth", // or get this from environment variables,
      signal: "", // if we don't have a signal, use the empty string
    };
    fetch(`https://developer.worldcoin.org/api/v1/verify/app_staging_09fa1cb7e6c0c60de6d8c56cd0dcb35c`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reqBody), 
    }).then((verifyRes) => {
      verifyRes.json().then((wldResponse) => {
        if (verifyRes.status == 200) {
          console.log(wldResponse);
          
          // this is where you should perform backend actions based on the verified credential
          // i.e. setting a user as "verified" in a database
          res.status(verifyRes.status).send({ code: "success" });
        } else {
          // return the error code and detail from the World ID /verify endpoint to our frontend
          res.status(verifyRes.status).send({ 
            code: wldResponse.code, 
            detail: wldResponse.detail 
          });
        }
      });
    });
  }
  