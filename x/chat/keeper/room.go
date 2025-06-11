package keeper

import (
	"aaronetwork/x/chat/types"
	"context"
	"log"
)

const (
	RoomHandlePrefix = "room-handle-"
	RoomIDPrefix     = "room-id-"
)

func (k Keeper) SetRoom(ctx context.Context, room types.Room) {
	store := k.storeService.OpenKVStore(ctx)

	bz, err := k.cdc.Marshal(&room)
	if err != nil {
		panic(err)
	}

	key := []byte(RoomIDPrefix + room.Id)
	if err := store.Set(key, bz); err != nil {
		log.Printf("Error when set room: %v", err)
		return
	}

	k.SetRoomHandleIndex(ctx, room.Handle, room.Id)
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

func (k Keeper) GetRoomByHandle(ctx context.Context, handle string) (types.Room, bool) {
	store := k.storeService.OpenKVStore(ctx)

	roomIDBytes, _ := store.Get([]byte(RoomHandlePrefix + handle))
	if roomIDBytes == nil {
		return types.Room{}, false
	}

	roomId := string(roomIDBytes)
	return k.GetRoom(ctx, roomId)
}

func (k Keeper) GetRoom(ctx context.Context, roomId string) (types.Room, bool) {
	store := k.storeService.OpenKVStore(ctx)

	roomBytes, _ := store.Get([]byte(RoomIDPrefix + roomId))
	if roomBytes == nil {
		return types.Room{}, false
	}

	var room types.Room
	err := room.Unmarshal(roomBytes)
	if err != nil {
		return types.Room{}, false
	}

	return room, true
}

func (k Keeper) SetRoomHandleIndex(ctx context.Context, handle, id string) {
	store := k.storeService.OpenKVStore(ctx)
	err := store.Set([]byte(RoomHandlePrefix+handle), []byte(id))
	if err != nil {
		log.Printf("Error when set room handle: %v", err)
		return
	}
}
