from behave import *
import requests


@given('we have behave installed')
def step_impl(context):
    pass

@given('todo')
def step_impl(context):
    print("Todo")

@when('we implement a test')
def step_impl(context):
    assert True is not False

@then('behave will test it for us!')
def step_impl(context):
    assert context.failed is False



@given('Manish')
def anc(context):
    print("Manish")

@when(u'Shweta')
def step_impl(context):
    print("Shweta")
    api_url = "https://jsonplaceholder.typicode.com/todos/1"
    context.resp = requests.get(api_url)
    print(context.resp.json())
    assert context.resp.status_code == 200

@then(u'Hitarth')
def step_impl(context):
    print("Hiti")
    print(context.resp)
    assert context.resp.status_code == 200

@then(u'Sushrut')
def step_impl(context):
    print("Sushi")
    print(context.resp.json())

@given('user enters name and password')
def step_impl(context):
    #access multiline text with .text attribute
    print("Multiline Text: " + context.text)

@then('user should be logged in')
def step_impl(context):
    print("Another ###############")