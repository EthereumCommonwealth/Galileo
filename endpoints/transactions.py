import falcon

from utils import web3_to_dict, convert_id_to_number_or_hash, Parity


class Transaction(object):
    __slots__ = ['web3']

    def __init__(self, web3):
        self.web3 = web3

    def on_get(self, req, resp, chain, transaction_hash):
        """
        Get Tx Details.

        Return a JSON

        :param chain: string - chain symbol
        :param transaction_hash: string - transaction hash
        :return:
        """
        web3 = self.web3[chain.upper()]

        transaction_hash = convert_id_to_number_or_hash(transaction_hash)

        transaction_details = web3.eth.getTransaction(transaction_hash)
        transaction_receipt = web3.eth.getTransactionReceipt(transaction_hash)

        if not transaction_details:
            resp.status = falcon.HTTP_404
            return

        transaction = {
            'transactionDetails': web3_to_dict(transaction_details),
            'transactionReceipt': web3_to_dict(transaction_receipt)
        }

        resp.media = transaction
        resp.status = falcon.HTTP_200
