cur = db.accounts.find({"accountId": {$exists: true}}, {accountId: 1, _id: 0});
while (cur.hasNext()) {
  acc = cur.next();
  print(acc.accountId);
}
