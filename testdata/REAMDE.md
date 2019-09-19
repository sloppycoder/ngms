## Generate test accounts data for benchmark

### install mimesis
```
pip install mimesis

```

### prep the data
```
# generate 100 test accounts 
python gen_accounts.py 100 > test_accounts.json

# import test data into database
mongoimport host/database < test_accounts.json

#create index 
 db.accounts.createIndex({accountId: 1})

# export account id into a text file to be used by load generator
mongo host/database --quiet account_ids.js > account_ids.txt

```