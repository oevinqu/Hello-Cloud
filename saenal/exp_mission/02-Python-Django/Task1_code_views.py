from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    html_content = """
    <html>
        <head>
            <title>Hello</title>
            <style>
                body {
                    justify-self: center;
                    align-content: center;
                    font-size: 75px;
                    font-weight: bold;
                }
            </style>
        </head>
        <body>
            <div>
                <p>Hello Django !!</p>
            </div>
        </body>
    </html>
    """
    return HttpResponse(html_content)
