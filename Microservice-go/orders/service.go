package main

import (
	"context"
	common "github.com/DeTiveNC/commons"
	pb "github.com/DeTiveNC/commons/api"
	"log"
)

type Service struct {
	store OrderStore
}

func NewService(store OrderStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateOrder(context.Context) error {
	return nil
}

func (s *Service) ValidateOder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergedItems := MergeItemsQuantity(p.Items)
	log.Println("Merged items: ", mergedItems)
	// validate stock service

	return nil
}

func MergeItemsQuantity(items []*pb.ItemsWithQuatity) []*pb.ItemsWithQuatity {
	merged := make([]*pb.ItemsWithQuatity, 0)

	for _, item := range items {
		found := false
		for _, mItem := range merged {
			if item.ID == mItem.ID {
				mItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
