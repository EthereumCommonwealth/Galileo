import falcon

from utils import web3_to_dict


class AccountInfo(object):
    __slots__ = ['web3']

    def __init__(self, web3):
        self.web3 = web3

    def on_get(self, req, resp, chain, address):
        """
        Get Account.

        Return a JSON

        :param chain: string - chain symbol
        :param address: string - account address 0x
        :return:
        """
        web3 = self.web3[chain.upper()]

        if not web3.isAddress(address):
            resp.status = falcon.HTTP_400
            return

        account = {
            'account': address,
            'balance': web3.eth.getBalance(address),
            'code': web3.eth.getCode(address),
            'transactionCount': web3.eth.getTransactionCount(address),
            'chain': chain
        }

        resp.media = web3_to_dict(account)
        resp.status = falcon.HTTP_200


class AccountBalance(object):
    __slots__ = ['web3']

    def __init__(self, web3):
        self.web3 = web3

    def on_get(self, req, resp, chain, address):
        """
        Get Account balance.

        Return a JSON

        :param chain: string - chain symbol
        :param address: string - account address 0x
        :return:
        """
        web3 = self.web3[chain.upper()]

        if not web3.isAddress(address):
            resp.status = falcon.HTTP_400
            return

        balance = web3.eth.getBalance(address)

        balance = {
            'account': address,
            'balance': balance,
            'chain': chain
        }

        resp.media = web3_to_dict(balance)
        resp.status = falcon.HTTP_200
