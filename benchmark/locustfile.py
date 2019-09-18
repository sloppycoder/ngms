from locust import HttpLocust, TaskSet, task

class InvokeAccountsApi(TaskSet):
    @task(1)
    def account_summary(self):
        self.client.get("/accounts/100-1234-5577-890")

class ApiStressTest(HttpLocust):
    task_set = InvokeAccountsApi
    min_wait = 100
    max_wait = 100
