./../natscli/nats stream add calvin -s localhost:6000 \
--max-msgs-per-subject=-1 \
--max-msg-size=-1 \
--max-msgs=-1 \
--max-bytes=-1 \
--max-age=-1 \
--replicas=1 \
--retention=limits \
--storage=file \
--subjects=calvin.sequencer \
--no-deny-delete \
--discard=old \
--dupe-window="2m0s" \
--no-allow-rollup \
--deny-purge  

./../natscli/nats consumer add \
-s localhost:6000 \
--pull \
--deliver=all \
--ack=explicit \
--replay=original \
--filter="" \
--max-deliver=-1 \
--max-pending=0 \
--no-headers-only \
--backoff=none \
calvin n0

./../natscli/nats consumer add \
-s localhost:6000 \
--pull \
--deliver=all \
--ack=explicit \
--replay=original \
--filter="" \
--max-deliver=-1 \
--max-pending=0 \
--no-headers-only \
--backoff=none \
calvin n1

./../natscli/nats consumer add \
-s localhost:6000 \
--pull \
--deliver=all \
--ack=explicit \
--replay=original \
--filter="" \
--max-deliver=-1 \
--max-pending=0 \
--no-headers-only \
--backoff=none \
calvin n2
