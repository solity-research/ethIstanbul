from os import environ
import logging
import requests
from transformers import pipeline,Conversation

logging.basicConfig(level="INFO")
logger = logging.getLogger(__name__)

rollup_server = environ["ROLLUP_HTTP_SERVER_URL"]
logger.info(f"HTTP rollup_server url is {rollup_server}")

pipe = pipeline("conversational",model="facebook/blenderbot-400M-distill")

def handle_advance(data):
    logger.info(f"Received advance request data {data}")
    logger.info("Changing data")
    data['payload'] = bytes.fromhex(data['payload'][2:]).decode('utf-8')
    conversation_1 = Conversation(data['payload'])
    result=pipe([conversation_1])
    ethereum_hex = "0x" + result.generated_responses[-1].encode("utf-8").hex()

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
