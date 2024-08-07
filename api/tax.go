package api

import (
	"context"
)

// Tax 个人所得税申报明细表
type Tax interface {
	// GetTaxFile 下载个人所得税申报明细表
	GetTaxFile(context.Context, *GetTaxFileRequest) (*GetTaxFileResponse, error)
	// GetUserCross 查询纳税人是否为跨集团用户
	GetUserCross(context.Context, *GetUserCrossRequest) (*GetUserCrossResponse, error)
}

// taxImpl Tax 接口实现
type taxImpl struct {
	cc Invoker
}

// NewTax 创建客户端
func NewTax(cc Invoker) Tax {
	return &taxImpl{cc}
}

// GetTaxFile 下载个人所得税申报明细表
func (c *taxImpl) GetTaxFile(ctx context.Context, in *GetTaxFileRequest) (*GetTaxFileResponse, error) {
	out := new(GetTaxFileResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/tax/v1/taxfile/download", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserCross 查询纳税人是否为跨集团用户
func (c *taxImpl) GetUserCross(ctx context.Context, in *GetUserCrossRequest) (*GetUserCrossResponse, error) {
	out := new(GetUserCrossResponse)
	err := c.cc.Invoke(ctx, "POST", "/api/tax/v1/user/cross", false, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTaxFileRequest 下载个人所得税申报明细表请求
type GetTaxFileRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 平台企业签约主体
	EntID string `json:"ent_id,omitempty"`
	// 所属期
	YearMonth string `json:"year_month,omitempty"`
}

// GetTaxFileResponse 下载个人所得税申报明细表返回
type GetTaxFileResponse struct {
	// 文件详情
	FileInfo []*FileInfo `json:"file_info,omitempty"`
}

// FileInfo 报税文件详情
type FileInfo struct {
	// 文件名称
	Name string `json:"name,omitempty"`
	// 下载文件临时 URL
	URL string `json:"url,omitempty"`
	// 文件解压缩密码
	Pwd string `json:"pwd,omitempty"`
}

// GetUserCrossRequest 查询纳税人是否为跨集团用户请求
type GetUserCrossRequest struct {
	// 平台企业 ID
	DealerID string `json:"dealer_id,omitempty"`
	// 年份
	Year string `json:"year,omitempty"`
	// 身份证号码
	IDCard string `json:"id_card,omitempty"`
	// 平台企业签约主体
	EntID string `json:"ent_id,omitempty"`
}

// GetUserCrossResponse 查询纳税人是否为跨集团用户返回
type GetUserCrossResponse struct {
	// 跨集团标识
	IsCross bool `json:"is_cross,omitempty"`
}
