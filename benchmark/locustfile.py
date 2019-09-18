from locust import HttpLocust, TaskSet, task

class InvokeRestApi(TaskSet):
    @task(1)
    def account_summary(self):
        self.client.get("/accounts/100-1234-5577-890")

class ApiStressTest(HttpLocust):
    task_set = InvokeRestApi
    min_wait = 10
    max_wait = 10
