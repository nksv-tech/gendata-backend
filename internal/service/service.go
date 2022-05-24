package service

import (
	"bytes"
	"context"
	"time"

	"github.com/google/wire"
	gendata "github.com/nikitaksv/gendata/pkg/service"
	v1 "github.com/nksv-tech/gendata-backend/pkg/gendata/v1"
	inmemcache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/durationpb"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGenDataService)

func NewGenDataService() *GenDataService {
	const cacheExpiration = 72 * time.Hour
	const cacheCleanInterval = 30 * time.Minute

	return &GenDataService{
		cache: inmemcache.New(cacheExpiration, cacheCleanInterval),
	}
}

type GenDataService struct {
	v1.UnimplementedGenDataServiceServer
	cache *inmemcache.Cache
}

func (s *GenDataService) Gen(ctx context.Context, request *v1.GenRequest) (*v1.GenResponse, error) {
	cacheKey := request.String()
	if get, ok := s.cache.Get(cacheKey); ok {
		return get.(*v1.GenResponse), nil
	}

	genreq := &gendata.GenRequest{
		Tmpl: request.GetTmpl(),
		Data: request.GetData(),
		Config: &gendata.Config{
			LangSettings: &gendata.LangSettings{
				ConfigMapping: &gendata.ConfigMapping{
					TypeMapping:      convertTypeMapping(request.GetConfig().GetLangSettings().GetConfigMapping().GetTypeMapping()),
					TypeDocMapping:   convertTypeMapping(request.GetConfig().GetLangSettings().GetConfigMapping().GetTypeDocMapping()),
					ClassNameMapping: request.GetConfig().GetLangSettings().GetConfigMapping().GetClassNameMapping(),
					FileNameMapping:  request.GetConfig().GetLangSettings().GetConfigMapping().GetFileNameMapping(),
				},
				Code:               request.GetConfig().GetLangSettings().GetCode(),
				Name:               request.GetConfig().GetLangSettings().GetName(),
				FileExtension:      request.GetConfig().GetLangSettings().GetFileExtension(),
				SplitObjectByFiles: request.GetConfig().GetLangSettings().GetSplitObjectByFiles(),
			},
			Lang:            request.GetConfig().GetLang(),
			DataFormat:      request.GetConfig().GetDataFormat(),
			RootClassName:   request.GetConfig().GetRootClassName(),
			PrefixClassName: request.GetConfig().GetPrefixClassName(),
			SuffixClassName: request.GetConfig().GetSuffixClassName(),
			SortProperties:  request.GetConfig().GetSortProperties(),
		},
	}
	gen, err := gendata.NewService(nil).Gen(ctx, genreq)
	if err != nil {
		return nil, errors.Wrapf(err, "Gen error")
	}

	rsp := &v1.GenResponse{
		RenderedFiles: make([]*v1.RenderedFile, len(gen.RenderedFiles)),
		RenderTime:    durationpb.New(gen.RenderTime),
	}
	for i, file := range gen.RenderedFiles {
		bs := bytes.Buffer{}
		_, err := bs.ReadFrom(file.Content)
		if err != nil {
			return nil, errors.Wrapf(err, "can't read from rendered file '%s' content", file.FileName)
		}
		rsp.RenderedFiles[i] = &v1.RenderedFile{
			Content:  bs.Bytes(),
			FileName: file.FileName,
		}
	}

	s.cache.SetDefault(cacheKey, rsp)

	return rsp, nil
}

func (s *GenDataService) PredefinedSettings(ctx context.Context, request *v1.PredefinedSettingsRequest) (*v1.PredefinedSettingsResponse, error) {
	list, err := gendata.NewService(nil).PredefinedLangSettings(ctx, &gendata.PredefinedLangSettingsListRequest{
		Code: request.GetCode(),
		Name: request.GetName(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "can't get list predifined languages settings")
	}

	items := make([]*v1.LangSettings, len(list.Items))
	for i, item := range list.Items {
		items[i] = &v1.LangSettings{
			Code:               item.Code,
			Name:               item.Name,
			FileExtension:      item.FileExtension,
			SplitObjectByFiles: item.SplitObjectByFiles,
			ConfigMapping: &v1.ConfigMapping{
				TypeMapping:      convertBackTypeMapping(item.ConfigMapping.TypeMapping),
				TypeDocMapping:   convertBackTypeMapping(item.ConfigMapping.TypeDocMapping),
				ClassNameMapping: item.ConfigMapping.ClassNameMapping,
				FileNameMapping:  item.ConfigMapping.FileNameMapping,
			},
		}
	}

	return &v1.PredefinedSettingsResponse{Items: items}, nil
}

func convertTypeMapping(tm *v1.TypeMapping) *gendata.TypeMapping {
	return &gendata.TypeMapping{
		Array:       tm.GetArray(),
		ArrayBool:   tm.GetArrayBool(),
		ArrayFloat:  tm.GetArrayFloat(),
		ArrayInt:    tm.GetArrayInt(),
		ArrayObject: tm.GetArrayObject(),
		ArrayString: tm.GetArrayString(),
		Bool:        tm.GetBool(),
		Float:       tm.GetFloat(),
		Int:         tm.GetInt(),
		Null:        tm.GetNull(),
		Object:      tm.GetObject(),
		String:      tm.GetString_(),
	}
}

func convertBackTypeMapping(tm *gendata.TypeMapping) *v1.TypeMapping {
	return &v1.TypeMapping{
		Array:       tm.Array,
		ArrayBool:   tm.ArrayBool,
		ArrayFloat:  tm.ArrayFloat,
		ArrayInt:    tm.ArrayInt,
		ArrayObject: tm.ArrayObject,
		ArrayString: tm.ArrayString,
		Bool:        tm.Bool,
		Float:       tm.Float,
		Int:         tm.Int,
		Null:        tm.Null,
		Object:      tm.Object,
		String_:     tm.String,
	}
}
