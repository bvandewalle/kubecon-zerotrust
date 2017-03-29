TRIREME_SECRET="trireme"
TRIREME_PSK_SECRET_ENTRY="triremepsk"

echo "Attempting to generate PSK"
PSK=$(openssl rand -base64 16)
kubectl create secret generic $TRIREME_SECRET --from-literal=$TRIREME_PSK_SECRET_ENTRY=$PSK
