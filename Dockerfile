FROM golang:1.8
ADD . /go/src/random-word-translator
RUN go install random-word-translator/word
#Usually not a good idea but this is a low threat API with no personal information attached
ENV TRANSLATION_API_KEY trnsl.1.1.20170513T204309Z.5861366a1aa3a27b.f3ebfce9182383782687fab1a4ae61432f3fa86b
EXPOSE 8080
