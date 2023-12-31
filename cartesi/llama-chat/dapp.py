from os import environ
import logging
import requests
from llama_cpp import Llama

logging.basicConfig(level="INFO")
logger = logging.getLogger(__name__)

rollup_server = "http://localhost:8080/host-runner"
logger.info(f"HTTP rollup_server url is {rollup_server}")
logger.info("Creating Llama instance")
llm = Llama(model_path="llama-2-7b-chat.Q4_K_M.gguf")
logger.info("Llama instance created")
protocol_data_dict = {}
def handle_advance(data):
    logger.info(f"Received advance request data {data}")
    data['payload'] = bytes.fromhex(data['payload'][2:]).decode('utf-8')
    logger.info(f"asking promt")
    print(data['payload'])
    output = llm(f"Q: {data['payload']} A: ", max_tokens=32, stop=["Q:", "\n"], echo=False)
    logger.info(f"output is {output['choices'][-1]['text']}")
    ethereum_hex = "0x" + output['choices'][-1]['text'].encode("utf-8").hex()
    logger.info("Adding notice")
    notice={'payload':ethereum_hex}
    response = requests.post(rollup_server + "/notice", json=notice)
    logger.info(f"Received notice status {response.status_code} body {response.content}")
    return "accept"


def handle_inspect(data):
    logger.info(f"Received inspect request data {data}")
    return "accept"


handlers = {
    "advance_state": handle_advance,
    "inspect_state": handle_inspect,
}

finish = {"status": "accept"}

while True:
    logger.info("Sending finish")
    response = requests.post(rollup_server + "/finish", json=finish)
    logger.info(f"Received finish status {response.status_code}")
    if response.status_code == 202:
        logger.info("No pending rollup request, trying again")
    else:
        rollup_request = response.json()
        data = rollup_request["data"]
        handler = handlers[rollup_request["request_type"]]
        finish["status"] = handler(rollup_request["data"])
