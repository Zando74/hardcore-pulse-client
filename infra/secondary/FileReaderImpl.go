package secondary

import (
	"deathlog-tracker/domain/entity"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type FileReaderImpl struct {}

func (f *FileReaderImpl) ExtractPlayerDeathLogData(filePath string) ([]entity.DeathRecord, error) {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(filePath); err != nil {
		return nil, fmt.Errorf("error loading Lua file: %w", err)
	}

	val := L.GetGlobal("deathlog_data")
	tbl, ok := val.(*lua.LTable)
	if !ok {
		return nil, fmt.Errorf("deathlog_data not found or invalid")
	}

	return flattenDeathlogTable(tbl), nil
}

func flattenDeathlogTable(tbl *lua.LTable) []entity.DeathRecord {
	var records []entity.DeathRecord

	tbl.ForEach(func(realmkey, realmVal lua.LValue) {
		if realmTbl, ok := realmVal.(*lua.LTable); ok {
			realmTbl.ForEach(func(_, playerVal lua.LValue) {
				if playerTbl, ok := playerVal.(*lua.LTable); ok {
					rec := entity.DeathRecord{}

					playerTbl.ForEach(func(k, v lua.LValue) {
						rec.Realm = realmkey.String()
						switch k.String() {
						case "class_id":
							rec.ClassID = toInt(v)
						case "guild":
							rec.Guild = v.String()
						case "date":
							rec.Date = toInt64(v)
						case "map_pos":
							rec.MapPos = v.String()
						case "source_id":
							rec.SourceID = toInt(v)
						case "name":
							rec.Name = v.String()
						case "last_words":
							rec.LastWords = v.String()
						case "level":
							rec.Level = toInt(v)
						case "map_id":
							rec.MapID = toInt(v)
						case "instance_id":
							rec.InstanceID = toInt(v)
						case "race_id":
							rec.RaceID = toInt(v)
						}
					})

					records = append(records, rec)
				}
			})
		}
	})

	return records
}

func toInt(v lua.LValue) int {
	if n, ok := v.(lua.LNumber); ok {
		return int(n)
	}
	return 0
}

func toInt64(v lua.LValue) int64 {
	if n, ok := v.(lua.LNumber); ok {
		return int64(n)
	}
	return 0
}