package service

import (
	"context"
	"fbsTest/internal/repository"
	"fmt"
	"log"
	"strings"
)

type FibService struct {
	repo repository.FibRepository
}

func NewFibService(repo repository.FibRepository) *FibService {
	return &FibService{
		repo: repo,
	}
}

func (f *FibService) GetFibSeq(ctx context.Context, start, end int) (string, error) {
	if start > end {
		return "", fmt.Errorf("start > end")
	}
	if start < 0 || end < 0 {
		return "", fmt.Errorf("start and end should be positive")
	}
	cachedSeq, err := f.repo.GetCachedFibSeq(ctx, end)
	if err != nil {
		log.Println("getCached error:", err)
	}
	if cachedSeq != nil {
		return f.seqFilter(cachedSeq, start, end), nil
	}

	seq := f.calcFibSeq(end)
	if err := f.repo.SetNewFibSeq(ctx, seq, end); err != nil {
		log.Println("save to redis error: ", err)
	}
	return f.seqFilter(seq, start, end), nil
}

func (f *FibService) calcFibSeq(n int) []int64 {
	fibSeq := make([]int64, n+1, n+2)
	if n < 2 {
		return []int64{0, 1}
	}
	fibSeq[0] = 0
	fibSeq[1] = 1
	for i := 2; i <= n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}
	return fibSeq
}

func (f *FibService) seqFilter(fibSeq []int64, start, end int) string {
	var strB strings.Builder

	for i := start; i <= end; i++ {
		_, _ = fmt.Fprintf(&strB, "%d-%d ", i, fibSeq[i])
	}

	res := strB.String()
	return res
}
