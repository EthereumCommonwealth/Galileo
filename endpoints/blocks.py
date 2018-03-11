import falcon

from utils import web3_to_dict, convert_id_to_number_or_hash


class BlockInfo(object):
    __slots__ = ['web3']

    def __init__(self, web3):
        self.web3 = web3

    def on_get(self, req, resp, chain, block_id):
        """
        Get Block Details.

        Return a JSON

        :param block_id: string - block number or hash
        :return:
        """
        web3 = self.web3[chain.upper()]

        block_id = convert_id_to_number_or_hash(block_id)

        block = web3.eth.getBlock(block_id, full_transactions=True)
        if not block:
            resp.status = falcon.HTTP_404
            return

        resp.media = web3_to_dict(block)
        resp.status = falcon.HTTP_200
