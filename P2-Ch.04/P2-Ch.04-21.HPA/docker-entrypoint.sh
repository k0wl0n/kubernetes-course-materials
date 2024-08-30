#!/bin/sh

if [ -z $NUM_THREADS ]
then
	echo "[!] Running the application with default thread"
       	openssl speed -multi 50 | tail /dev/zero
else
	echo "[*] Running the application with number of thread: " $NUM_THREADS
	openssl speed -multi $NUM_THREADS | tail /dev/zero
fi	