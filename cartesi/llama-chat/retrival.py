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
def savemessage(input_data):
    protocol_data_dict[input_data[1]]=input_data[2]

def returnanswer(input_data):
    context = protocol_data_dict[input_data[1]]
    question =input_data[2]
    input_text = f"Context: {context} Question: {question} Answer:"
    answer = llm(input_text,max_tokens=256,
    temperature=0.1,
    top_p=0.5,
    echo=False,
    stop=["#"])
    logger.info(f"output is {answer['choices'][-1]['text']}")

    return answer['choices'][-1]['text']
def handle_advance(data):
    logger.info(f"Received advance request data {data}")
    data['payload'] = bytes.fromhex(data['payload'][2:]).decode('utf-8')
    logger.info(f"asking promt")
    print(data['payload'])
    input_data=data['payload'].split(":")
    if input_data[0] == "crt":
        if input_data[1] in protocol_data_dict.keys():
            "0x" + f"{input_data[1]} is exist".encode("utf-8").hex()
        else:
            savemessage(input_data)
            logger.info(f"{input_data[1]} data is saved")
            ethereum_hex = "0x" + f"{input_data[1]} is added".encode("utf-8").hex()
    elif input_data[0]=='msg':
        answer=returnanswer(input_data)
        logger.info(f"Model Response: {answer}")
        ethereum_hex = "0x" + f"{answer}".encode("utf-8").hex()
    else:
        logger.log('there is problem')
        ethereum_hex = "0x" + f"there is problem".encode("utf-8").hex()


    logger.info("Adding notice")
    notice={'payload':ethereum_hex}
    response = requests.post(rollup_server + "/notice", json=notice)
    logger.info(f"Received notice status {response.status_code} body {response.content}")
    return "accept"


def handle_inspect(data):
    logger.info(f"Received inspect request data {data}")
    report = {"payload":  protocol_data_dict[data]}
    response = requests.post(rollup_server + "/report", json=report)
    logger.info(f"Received report status {response.status_code}")
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
