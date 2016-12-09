import requests

session = requests.session()

def Login():
    data = {
        "username" : "kevince",
        "password": "abc",
    }
    r = session.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
    print(r.json())

def Submit():
    data = {
        "problem_sid" : "zoj#1000",
        "code":"""
    #include <iostream>
    using namespace std;
    int main() {
        int a, b;
        cin >> a >> b;
        cout << a + b << endl;
        return 0;
    }
        """,
        "language_id": 1,
        "is_shared" : False,

    }

    r = session.post("http://127.0.0.1:8000/api/inline/submit", json=data)
    print(r.status_code)
    print(r.json())

if __name__ == '__main__':
    Login()
    Submit()