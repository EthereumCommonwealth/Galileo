import falcon

from web3 import Web3, IPCProvider, HTTPProvider

from endpoints.accounts import AccountBalance, AccountInfo
from endpoints.blocks import BlockInfo
from endpoints.transactions import Transaction

web3 = {
    'CLO': Web3(IPCProvider()),
}

api = falcon.API()
api.add_route('/{chain}/block/{block_id}', BlockInfo(web3))
api.add_route('/{chain}/transaction/{transaction_hash}', Transaction(web3))
api.add_route('/{chain}/account/{address}', AccountInfo(web3))
api.add_route('/{chain}/account/{address}/balance', AccountBalance(web3))
api.add_route('/{chain}/account/{address}/transaction', None)

