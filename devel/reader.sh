firecore start reader-node --config-file= --log-format=stackdriver --log-to-file=false --data-dir=data --common-auto-max-procs --common-auto-mem-limit-percent=90 --common-one-block-store-url=data/oneblock --common-first-streamable-block=0 --reader-node-data-dir=data/oneblock --reader-node-working-dir=data/work --reader-node-readiness-max-latency=600s --reader-node-debug-firehose-logs=false --reader-node-blocks-chan-capacity=1000 --reader-node-grpc-listen-addr=:9001 --reader-node-manager-api-addr=:8080 --reader-node-path=firevara --reader-node-arguments='fetch 0 --state-dir data --endpoints https://vara-mainnet.public.blastapi.io'
