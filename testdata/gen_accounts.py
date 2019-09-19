import json, re, sys
from mimesis import Person, random

PRODUCTS = [
    [ '1001', 'Super Saver Account' ],
    [ '1002', 'Salary Plus Account' ],
    [ '2001', 'Student Saver Account' ],
    [ '2002', 'Everyday Rebate Account']
]

NUM_OF_PRODUCTS = len(PRODUCTS)

all_account_ids = {}


def unique_account_id(p):
    # try generate a unique account number for 10 times
    # use telephone to proxy an account number
    for i in range(10):
        num = re.sub(r"\+|\-|\(|\)|\.|\s", '', p.telephone())
        if num not in all_account_ids:
            all_account_ids[num] = 1
            return num
    return None


def gen_prod(rand):
    prod_code, prod_name  = PRODUCTS[int(rand.uniform(0,NUM_OF_PRODUCTS-1,0))]
    return prod_code, prod_name


def gen_account(n):
    p = Person(locales.EN)
    r = random.Random()
    for i in range(n):
        account_id = unique_account_id(p)
        prod_code, prod_name = gen_prod(r)
        balance = r.uniform(100.0, 1_000_000.0, 4)
        account = {
            'accountId': account_id,
            'prodCode': prod_code,
            'prodName': prod_name,
            'balances': [
                {'amount': balance, 'type': '0', 'creditFlag': False},
                {'amount': balance, 'type': '1', 'creditFlag': False}
            ]
        }
        print(json.dumps(account))


if __name__ == '__main__':
    iter = 1
    if len(sys.argv) > 1:
        try:
            iter = int(sys.argv[1])
        except ValueError as e:
            pass

    gen_account(iter)
