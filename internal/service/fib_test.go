package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeqFilter(t *testing.T) {
	res := "3-2 4-3 5-5 6-8 7-13 "
	fibSeq := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	gotRes := testService.seqFilter(fibSeq, 3, 7)
	require.Equal(t, res, gotRes)
}

func TestCalcFibSeq(t *testing.T) {
	fibSeq := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	gotFibSeq := testService.calcFibSeq(9)
	require.Equal(t, fibSeq, gotFibSeq)
}

func TestGetFibSeq(t *testing.T) {
	res := "3-2 4-3 5-5 6-8 7-13 "

	ctx := context.Background()
	_, err := testService.GetFibSeq(ctx, 5, 0)
	require.Error(t, err)
	_, err = testService.GetFibSeq(ctx, -1, 4)
	require.Error(t, err)
	_, err = testService.GetFibSeq(ctx, 5, -10)
	require.Error(t, err)
	_, err = testService.GetFibSeq(ctx, -1, -4)
	require.Error(t, err)

	gotRes, err := testService.GetFibSeq(ctx, 3, 7)
	require.NoError(t, err)
	require.Equal(t, gotRes, res)
}
