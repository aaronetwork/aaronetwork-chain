package keeper

import (
	"aaronetwork/x/chat/types"
	"context"
)

func (k Keeper) SetRoom(ctx context.Context, room types.Room) {
	store := k.storeService.OpenKVStore(ctx)

	bz, err := k.cdc.Marshal(&room)
	if err != nil {
		panic(err)
	}

	key := []byte("room:" + room.Id)
	if err := store.Set(key, bz); err != nil {
		panic(err)
	}
}

func (k Keeper) GetAllRoom(ctx context.Context) []types.Room {
	store := k.storeService.OpenKVStore(ctx)
	iterator, _ := store.Iterator(nil, nil)
	defer iterator.Close()

	var rooms []types.Room
	for ; iterator.Valid(); iterator.Next() {
		var room types.Room
		err := k.cdc.Unmarshal(iterator.Value(), &room)
		if err != nil {
			continue
		}
		rooms = append(rooms, room)
	}

	return rooms
}
