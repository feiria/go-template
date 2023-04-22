#!/bin/sh

DATABASE=${1}
OUTDIR=${2}

# shellcheck disable=SC2086
gormt -d ${DATABASE} -o ${OUTDIR}
