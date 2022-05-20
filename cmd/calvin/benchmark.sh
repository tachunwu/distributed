./../natscli/nats bench calvin --js --pub 3 --size 16 --msgs 1000000 --storage=file -s localhost:6000 &
./../natscli/nats bench calvin --js --sub 3 --msgs 1000000 --storage=file -s localhost:6000