from hexbytes import HexBytes
from web3.module import Module
from web3.utils.datastructures import AttributeDict


class Parity(Module):
    def parityFilter(self, params_object):
        """
        parity = Parity(web3)
        txs = parity.parityFilter({'fromBlock': web3.toHex(5000), 'fromAddress': ['0x4a72b92cc0930f88c18ccb1bbf63d4567315d7c0']})

        :param params_object:
        :return:
        """
        return self.web3.manager.request_blocking(
                "trace_filter",
                [params_object],
        )


def web3_to_dict(object):
    new_object = {}

    for key in object.keys():
        value = object[key]

        if isinstance(value, HexBytes):
            value = value.hex()

        if isinstance(value, list):
            new_value = []
            for v in value:
                if isinstance(v, AttributeDict):
                    new_value.append(web3_to_dict(v))
                else:
                    new_value.append(v)
            value = new_value

        new_object[key] = value

    return new_object


def convert_id_to_number_or_hash(block_id):
    try:
        return int(block_id)
    except ValueError:
        return HexBytes(block_id)
