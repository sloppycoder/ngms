import datetime, json, re, sys, time
from mimesis import Generic, random

PRODUCTS = [
    [ '1001', 'Super Saver Account' ],
    [ '1002', 'Salary Plus Account' ],
    [ '2001', 'Student Saver Account' ],
    [ '2002', 'Everyday Rebate Account']
]

NUM_OF_PRODUCTS = len(PRODUCTS)

all_account_ids = {}


def unique_account_id(g):
    # try generate a unique account number for 10 times
    # use telephone to proxy an account number
    for i in range(10):
        if i > 1 :
            num = re.sub(r"\+|\-|\(|\)|\.|\s", '', g.person.telephone())
        else:
            # always use this account number for 1st account
            num = '111'

        if num not in all_account_ids:
            all_account_ids[num] = 1
            return num
    return None


def gen_prod(rand):
    prod_code, prod_name  = PRODUCTS[int(rand.uniform(0,NUM_OF_PRODUCTS-1,0))]
    return prod_code, prod_name


def gen_account(n):
    g = Generic('en')
    r = random.Random()
    ts_base = int(time.mktime(datetime.date(2019, 1, 1).timetuple()))

    for i in range(n):
        account_id = unique_account_id(g)
        prod_code, prod_name = gen_prod(r)
        balance = r.uniform(1000.0, 5_000_000.0, 4)
        ts = int(r.uniform(0, 15_724_800, 0)) # 3600 * 24 * 182
        tstr = {'$date': {'$numberLong': str((ts_base +ts)*1000)}}
        account = {
            'accountId': account_id,
            'nickName': g.person.name(), # use person name to proxy nickname
            'prodCode': prod_code,
            'prodName': prod_name,
            'currency': 'THB',
            'servicer': g.code.pin(), # 4 digits
            'status' : '0',
            'status_last_updated': tstr,
            'balances': [
                {'amount': balance, 'type': '0', 'creditFlag': False, 'last_updated': tstr},
                {'amount': balance, 'type': '1', 'creditFlag': False, 'last_updated': tstr}
            ]
        }
        if n < 5 and sys.stdout.isatty():
            print(json.dumps(account, indent=4, sort_keys=True))
        else:
            print(json.dumps(account))


if __name__ == '__main__':
    iter = 1
    if len(sys.argv) > 1:
        try:
            iter = int(sys.argv[1])
        except ValueError as e:
            pass

    gen_account(iter)
