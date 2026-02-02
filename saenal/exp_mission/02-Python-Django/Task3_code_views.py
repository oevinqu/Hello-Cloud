from django.shortcuts import render
from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.response import Response
from .serializers import ArticleSerializer
from .models import Article

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

class ArticleView(APIView):
    def get(self, request):
        articles = Article.objects.all()
        serializer = ArticleSerializer(articles, many=True)
        return Response(serializer.data)

    def post(self, request):
        serializer = ArticleSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=201)
        return Response(serializer.errors, status=400)
