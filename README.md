# AI-Driven Insights on Blockchain

Our platform uniquely combines AI with blockchain technology, offering an innovative approach to GPT models specifically tailored for various protocols.

### Key Use Cases

#### 1. **Marketplace Exploration**
- **Empower Community-Driven Innovation**: In our marketplace, we  provide a platform for users to contribute their own custom, protocol-oriented GPTs. This approach enriches our marketplace, fostering a community-driven ecosystem where innovation thrives.
- **Protocol-Specific Knowledge**: Each model is an expert in its protocol, offering deep technical and code-level insights.

#### 2. **Interactive Customization and Learning**
- **Engage with Protocol GPTs**: Interact with models like the Aave V3 Protocol GPT for an in-depth understanding of the protocol's technical and code aspects.
- **Tailored Learning Experience**: These models serve as intelligent guides, helping users grasp complex protocol functionalities.

### Powered by Cartesi

- **Scalable and Efficient**: Leveraging Cartesi's Linux-based virtual machine for scalable and efficient DApp development.
- **Fully Verifiable**: Ensuring full verification of GPT models, their training datasets, and model parameters.

### Transparency and Verification

- **On-Chain Transparency**: Our approach ensures that all aspects of model architecture and data usage are transparent and verifiable on the blockchain.
- **Credibility and Trust**: Sets a new standard in AI and blockchain integration, focusing on trust and credibility within our ecosystem.

## How It Works?

We have a login mechanism that takes wallet address and then shows WorldCoin QR. When a user finishes the authentication, the json result is sent to our smart contract to verify user on chain. After the verification whenever the user want to create a protocol, if he/she is verified, the creation process is started. 

In order to validate users uniqueness we use Proof of personhood of WorldCoin. In our first smart contract that interacts with our frontend we receive the root, nullifierHash and the proof which was generated on the frontend. then we use them to verify that this user is unique and interacts with us through our frontend. after verification then we let the user send a prompt to our personalised GPT Model. in Cartesi. This happens through Hyperlane. Our verifier contract sends the prompt or the GPTData contract creation method via Hyperlane to our another contract in Scroll. Scroll provides proofs for each tx thanks to its zk-evm nature so every step until our LLM model is verifiable. Then those messages come to our Factory Contract in Scroll and according to commands either a new data contract is created to feed our model or we directly send our prompt to our model.

We used huggingface models in our first trial. Due to Cartesi Linux runtime uses risc-v architecture some of the dependencies cant be compiled. Because of that we used LLAMA 2 cpp version. With this version ve are using llama 2 7B version with 4 bit quantization. So we are using less memory for loading model. This approach solved our dependency issue. Hence our current DApp which serves as custom gpt model almost ready to deploy Cartesi Machine We are also using inspect requests in our dapp to in order to show what information we stored so we are ensuring transparency.

## Project Structure

The ProtocolGPT project is organized as follows:

- `cartesi/`: Contains the example Cartesi Dapp, example transformer base GPT and our main LLM (cpp) model .
- `contracts/`: Contains all the Solidity codes which are used to tranfer the user prompts on chian.`.
- `frontend-worldcoin/`: The Next.JS frontend for the application. Run with the command npm run dev.
- 
## Deployed Contracts

| Name          | Address                                    | Chain    |
|---------------|--------------------------------------------|----------|
| WorldCoinVerifyer | [0x5E622554E119C8Dc90b4BE2ae2fF9D4fA034645C](https://mumbai.polygonscan.com/address/0x5E622554E119C8Dc90b4BE2ae2fF9D4fA034645C) | Polygon Mumbai |
| ProtocolGPTFactory | [0xebfE3C8dBc94ae65dD0EBB27A5687967b94Cf093](https://sepolia-blockscout.scroll.io/address/0xebfE3C8dBc94ae65dD0EBB27A5687967b94Cf093/internal-transactions#address-tabs) | Scroll Sepholia  |
| ProtocolGPTToken | [0x650F547C84b12458186c002e5Df58b9cDB1F23C0](https://scan.chiliz.com/address/0x650F547C84b12458186c002e5Df58b9cDB1F23C0) | Chiliz Mainnet  |


