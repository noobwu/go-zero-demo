// Code generated by goctl. DO NOT EDIT.
package types

type DeleteCartItemReq struct {
	Id int64 `json:"id"`
}

type DeleteCartItemResp struct {
}

type DeleteOrderReq struct {
	Id int64 `json:"id"`
}

type DeleteOrderResp struct {
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}

type ListCartItemReq struct {
	Current  int64 `form:"current"`
	PageSize int64 `form:"pageSize"`
}

type ListCartItemResp struct {
	Current  int64                `json:"current"`
	Data     []*ListtCartItemData `json:"data"`
	PageSize int64                `json:"pageSize"`
	Success  bool                 `json:"success"`
	Total    int64                `json:"total"`
}

type ListHomeDisplayReq struct {
	Current  int64 `form:"current"`
	PageSize int64 `form:"pageSize"`
}

type ListHomeDisplayResp struct {
	Current  int64                   `json:"current"`
	Data     []*ListtHomeDisplayData `json:"data"`
	PageSize int64                   `json:"pageSize"`
	Success  bool                    `json:"success"`
	Total    int64                   `json:"total"`
}

type ListOrderReq struct {
	Current  int64 `form:"current"`
	PageSize int64 `form:"pageSize"`
}

type ListOrderResp struct {
	Current  int64             `json:"current"`
	Data     []*ListtOrderData `json:"data"`
	PageSize int64             `json:"pageSize"`
	Success  bool              `json:"success"`
	Total    int64             `json:"total"`
}

type ListProductCategoryReq struct {
	Current  int64 `form:"current"`
	PageSize int64 `form:"pageSize"`
}

type ListProductCategoryResp struct {
	Current  int64                       `json:"current"`
	Data     []*ListtProductCategoryData `json:"data"`
	PageSize int64                       `json:"pageSize"`
	Success  bool                        `json:"success"`
	Total    int64                       `json:"total"`
}

type ListProductReq struct {
	Current  int64 `form:"current"`
	PageSize int64 `form:"pageSize"`
}

type ListProductResp struct {
	Current  int64               `json:"current"`
	Data     []*ListtProductData `json:"data"`
	PageSize int64               `json:"pageSize"`
	Success  bool                `json:"success"`
	Total    int64               `json:"total"`
}

type ListtCartItemData struct {
	Id                int64   `json:"id"`
	ProductId         int64   `json:"product_id"`
	ProductSkuId      int64   `json:"product_sku_id"`
	MemberId          int64   `json:"member_id"`
	Quantity          int64   `json:"quantity"`            // 购买数量
	Price             float64 `json:"price"`               // 添加到购物车的价格
	ProductPic        string  `json:"product_pic"`         // 商品主图
	ProductName       string  `json:"product_name"`        // 商品名称
	ProductSubTitle   string  `json:"product_sub_title"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"product_sku_code"`    // 商品sku条码
	MemberNickname    string  `json:"member_nickname"`     // 会员昵称
	CreateDate        string  `json:"create_date"`         // 创建时间
	ModifyDate        string  `json:"modify_date"`         // 修改时间
	DeleteStatus      int64   `json:"delete_status"`       // 是否删除
	ProductCategoryId int64   `json:"product_category_id"` // 商品分类
	ProductBrand      string  `json:"product_brand"`
	ProductSn         string  `json:"product_sn"`
	ProductAttr       string  `json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type ListtHomeDisplayData struct {
	Id              int64  `json:"id"`
	BrandId         int64  `json:"brand_id"`
	BrandName       string `json:"brand_name"`
	RecommendStatus int64  `json:"recommend_status"`
	Sort            int64  `json:"sort"`
}

type ListtOrderData struct {
	Id                    int64   `json:"id"` // 订单id
	MemberId              int64   `json:"member_id"`
	CouponId              int64   `json:"coupon_id"`
	OrderSn               string  `json:"order_sn"`                // 订单编号
	CreateTime            string  `json:"create_time"`             // 提交时间
	MemberUsername        string  `json:"member_username"`         // 用户帐号
	TotalAmount           float64 `json:"total_amount"`            // 订单总金额
	PayAmount             float64 `json:"pay_amount"`              // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freight_amount"`          // 运费金额
	PromotionAmount       float64 `json:"promotion_amount"`        // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64 `json:"integration_amount"`      // 积分抵扣金额
	CouponAmount          float64 `json:"coupon_amount"`           // 优惠券抵扣金额
	DiscountAmount        float64 `json:"discount_amount"`         // 管理员后台调整订单使用的折扣金额
	PayType               int64   `json:"pay_type"`                // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int64   `json:"source_type"`             // 订单来源：0->PC订单；1->app订单
	Status                int64   `json:"status"`                  // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int64   `json:"order_type"`              // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string  `json:"delivery_company"`        // 物流公司(配送方式)
	DeliverySn            string  `json:"delivery_sn"`             // 物流单号
	AutoConfirmDay        int64   `json:"auto_confirm_day"`        // 自动确认时间（天）
	Integration           int64   `json:"integration"`             // 可以获得的积分
	Growth                int64   `json:"growth"`                  // 可以活动的成长值
	PromotionInfo         string  `json:"promotion_info"`          // 活动信息
	BillType              int64   `json:"bill_type"`               // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string  `json:"bill_header"`             // 发票抬头
	BillContent           string  `json:"bill_content"`            // 发票内容
	BillReceiverPhone     string  `json:"bill_receiver_phone"`     // 收票人电话
	BillReceiverEmail     string  `json:"bill_receiver_email"`     // 收票人邮箱
	ReceiverName          string  `json:"receiver_name"`           // 收货人姓名
	ReceiverPhone         string  `json:"receiver_phone"`          // 收货人电话
	ReceiverPostCode      string  `json:"receiver_post_code"`      // 收货人邮编
	ReceiverProvince      string  `json:"receiver_province"`       // 省份/直辖市
	ReceiverCity          string  `json:"receiver_city"`           // 城市
	ReceiverRegion        string  `json:"receiver_region"`         // 区
	ReceiverDetailAddress string  `json:"receiver_detail_address"` // 详细地址
	Note                  string  `json:"note"`                    // 订单备注
	ConfirmStatus         int64   `json:"confirm_status"`          // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int64   `json:"delete_status"`           // 删除状态：0->未删除；1->已删除
	UseIntegration        int64   `json:"use_integration"`         // 下单时使用的积分
	PaymentTime           string  `json:"payment_time"`            // 支付时间
	DeliveryTime          string  `json:"delivery_time"`           // 发货时间
	ReceiveTime           string  `json:"receive_time"`            // 确认收货时间
	CommentTime           string  `json:"comment_time"`            // 评价时间
	ModifyTime            string  `json:"modify_time"`             // 修改时间
}

type ListtProductCategoryData struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parent_id"` // 上机分类的编号：0表示一级分类
	Name         string `json:"name"`
	Level        int64  `json:"level"` // 分类级别：0->1级；1->2级
	ProductCount int64  `json:"product_count"`
	ProductUnit  string `json:"product_unit"`
	NavStatus    int64  `json:"nav_status"`  // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int64  `json:"show_status"` // 显示状态：0->不显示；1->显示
	Sort         int64  `json:"sort"`
	Icon         string `json:"icon"` // 图标
	Keywords     string `json:"keywords"`
	Description  string `json:"description"` // 描述
}

type ListtProductData struct {
	Id                         int64   `json:"id"`
	BrandId                    int64   `json:"brand_id"`
	ProductCategoryId          int64   `json:"product_category_id"`
	FeightTemplateId           int64   `json:"feight_template_id"`
	ProductAttributeCategoryId int64   `json:"product_attribute_category_id"`
	Name                       string  `json:"name"`
	Pic                        string  `json:"pic"`
	ProductSn                  string  `json:"product_sn"`       // 货号
	DeleteStatus               int64   `json:"delete_status"`    // 删除状态：0->未删除；1->已删除
	PublishStatus              int64   `json:"publish_status"`   // 上架状态：0->下架；1->上架
	NewStatus                  int64   `json:"new_status"`       // 新品状态:0->不是新品；1->新品
	RecommandStatus            int64   `json:"recommand_status"` // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int64   `json:"verify_status"`    // 审核状态：0->未审核；1->审核通过
	Sort                       int64   `json:"sort"`             // 排序
	Sale                       int64   `json:"sale"`             // 销量
	Price                      float64 `json:"price"`
	PromotionPrice             float64 `json:"promotion_price"` // 促销价格
	GiftGrowth                 int64   `json:"gift_growth"`     // 赠送的成长值
	GiftPoint                  int64   `json:"gift_point"`      // 赠送的积分
	UsePointLimit              int64   `json:"use_point_limit"` // 限制使用的积分数
	SubTitle                   string  `json:"sub_title"`       // 副标题
	Description                string  `json:"description"`     // 商品描述
	OriginalPrice              float64 `json:"original_price"`  // 市场价
	Stock                      int64   `json:"stock"`           // 库存
	LowStock                   int64   `json:"low_stock"`       // 库存预警值
	Unit                       string  `json:"unit"`            // 单位
	Weight                     float64 `json:"weight"`          // 商品重量，默认为克
	PreviewStatus              int64   `json:"preview_status"`  // 是否为预告商品：0->不是；1->是
	ServiceIds                 string  `json:"service_ids"`     // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string  `json:"keywords"`
	Note                       string  `json:"note"`
	AlbumPics                  string  `json:"album_pics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string  `json:"detail_title"`
	DetailDesc                 string  `json:"detail_desc"`
	DetailHtml                 string  `json:"detail_html"`           // 产品详情网页内容
	DetailMobileHtml           string  `json:"detail_mobile_html"`    // 移动端网页详情
	PromotionStartTime         string  `json:"promotion_start_time"`  // 促销开始时间
	PromotionEndTime           string  `json:"promotion_end_time"`    // 促销结束时间
	PromotionPerLimit          int64   `json:"promotion_per_limit"`   // 活动限购数量
	PromotionType              int64   `json:"promotion_type"`        // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string  `json:"brand_name"`            // 品牌名称
	ProductCategoryName        string  `json:"product_category_name"` // 商品分类名称
}

type MemberLoginReq struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type MemberLoginResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	JwtToken
}

type MemberRegisterReq struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type MemberRegisterResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	JwtToken
}

type OrderQueryReq struct {
	BusinessId string `json:"businessId"` // 业务编号
	MerId      string `json:"merId"`      // 本地商户id
	PayType    string `json:"payType"`
}

type OrderQueryResp struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	PayMessage string `json:"payMessage"`
	PayStatus  string `json:"payStatus"` //0：初始化 1：已发送 2：成功 3：失败
}

type UnifiedOrderReq struct {
	BusinessId string `json:"businessId"`    // 业务编号
	Amount     string `json:"amount"`        // 金额
	PayType    string `json:"payType"`       // 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
	Remarks    string `json:"remarks"`       // 描述
	Code       string `json:"code,optional"` // jsapi支付的时候,需要根据临时code来获取openid
	MerId      string `json:"merId"`         // 本地商户id
}

type UnifiedOrderResp struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type UpdateCartItemReq struct {
	Id                int64   `json:"id"`
	ProductId         int64   `json:"product_id"`
	ProductSkuId      int64   `json:"product_sku_id"`
	MemberId          int64   `json:"member_id"`
	Quantity          int64   `json:"quantity"`            // 购买数量
	Price             float64 `json:"price"`               // 添加到购物车的价格
	ProductPic        string  `json:"product_pic"`         // 商品主图
	ProductName       string  `json:"product_name"`        // 商品名称
	ProductSubTitle   string  `json:"product_sub_title"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"product_sku_code"`    // 商品sku条码
	MemberNickname    string  `json:"member_nickname"`     // 会员昵称
	DeleteStatus      int64   `json:"delete_status"`       // 是否删除
	ProductCategoryId int64   `json:"product_category_id"` // 商品分类
	ProductBrand      string  `json:"product_brand"`
	ProductSn         string  `json:"product_sn"`
	ProductAttr       string  `json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type UpdateCartItemResp struct {
}

type UpdateOrderReq struct {
	Id                    int64   `json:"id"` // 订单id
	MemberId              int64   `json:"member_id"`
	CouponId              int64   `json:"coupon_id"`
	OrderSn               string  `json:"order_sn"`                // 订单编号
	CreateTime            string  `json:"create_time"`             // 提交时间
	MemberUsername        string  `json:"member_username"`         // 用户帐号
	TotalAmount           float64 `json:"total_amount"`            // 订单总金额
	PayAmount             float64 `json:"pay_amount"`              // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freight_amount"`          // 运费金额
	PromotionAmount       float64 `json:"promotion_amount"`        // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64 `json:"integration_amount"`      // 积分抵扣金额
	CouponAmount          float64 `json:"coupon_amount"`           // 优惠券抵扣金额
	DiscountAmount        float64 `json:"discount_amount"`         // 管理员后台调整订单使用的折扣金额
	PayType               int64   `json:"pay_type"`                // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int64   `json:"source_type"`             // 订单来源：0->PC订单；1->app订单
	Status                int64   `json:"status"`                  // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int64   `json:"order_type"`              // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string  `json:"delivery_company"`        // 物流公司(配送方式)
	DeliverySn            string  `json:"delivery_sn"`             // 物流单号
	AutoConfirmDay        int64   `json:"auto_confirm_day"`        // 自动确认时间（天）
	Integration           int64   `json:"integration"`             // 可以获得的积分
	Growth                int64   `json:"growth"`                  // 可以活动的成长值
	PromotionInfo         string  `json:"promotion_info"`          // 活动信息
	BillType              int64   `json:"bill_type"`               // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string  `json:"bill_header"`             // 发票抬头
	BillContent           string  `json:"bill_content"`            // 发票内容
	BillReceiverPhone     string  `json:"bill_receiver_phone"`     // 收票人电话
	BillReceiverEmail     string  `json:"bill_receiver_email"`     // 收票人邮箱
	ReceiverName          string  `json:"receiver_name"`           // 收货人姓名
	ReceiverPhone         string  `json:"receiver_phone"`          // 收货人电话
	ReceiverPostCode      string  `json:"receiver_post_code"`      // 收货人邮编
	ReceiverProvince      string  `json:"receiver_province"`       // 省份/直辖市
	ReceiverCity          string  `json:"receiver_city"`           // 城市
	ReceiverRegion        string  `json:"receiver_region"`         // 区
	ReceiverDetailAddress string  `json:"receiver_detail_address"` // 详细地址
	Note                  string  `json:"note"`                    // 订单备注
	ConfirmStatus         int64   `json:"confirm_status"`          // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int64   `json:"delete_status"`           // 删除状态：0->未删除；1->已删除
	UseIntegration        int64   `json:"use_integration"`         // 下单时使用的积分
}

type UpdateOrderResp struct {
}

type AddCartItemReq struct {
	ProductId         int64   `json:"product_id"`
	ProductSkuId      int64   `json:"product_sku_id"`
	MemberId          int64   `json:"member_id"`
	Quantity          int64   `json:"quantity"`            // 购买数量
	Price             float64 `json:"price"`               // 添加到购物车的价格
	ProductPic        string  `json:"product_pic"`         // 商品主图
	ProductName       string  `json:"product_name"`        // 商品名称
	ProductSubTitle   string  `json:"product_sub_title"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"product_sku_code"`    // 商品sku条码
	MemberNickname    string  `json:"member_nickname"`     // 会员昵称
	DeleteStatus      int64   `json:"delete_status"`       // 是否删除
	ProductCategoryId int64   `json:"product_category_id"` // 商品分类
	ProductBrand      string  `json:"product_brand"`
	ProductSn         string  `json:"product_sn"`
	ProductAttr       string  `json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type AddCartItemResp struct {
}

type AddOrderReq struct {
	MemberId              int64   `json:"member_id"`
	CouponId              int64   `json:"coupon_id"`
	OrderSn               string  `json:"order_sn"`                // 订单编号
	CreateTime            string  `json:"create_time"`             // 提交时间
	MemberUsername        string  `json:"member_username"`         // 用户帐号
	TotalAmount           float64 `json:"total_amount"`            // 订单总金额
	PayAmount             float64 `json:"pay_amount"`              // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freight_amount"`          // 运费金额
	PromotionAmount       float64 `json:"promotion_amount"`        // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64 `json:"integration_amount"`      // 积分抵扣金额
	CouponAmount          float64 `json:"coupon_amount"`           // 优惠券抵扣金额
	DiscountAmount        float64 `json:"discount_amount"`         // 管理员后台调整订单使用的折扣金额
	PayType               int64   `json:"pay_type"`                // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int64   `json:"source_type"`             // 订单来源：0->PC订单；1->app订单
	Status                int64   `json:"status"`                  // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int64   `json:"order_type"`              // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string  `json:"delivery_company"`        // 物流公司(配送方式)
	DeliverySn            string  `json:"delivery_sn"`             // 物流单号
	AutoConfirmDay        int64   `json:"auto_confirm_day"`        // 自动确认时间（天）
	Integration           int64   `json:"integration"`             // 可以获得的积分
	Growth                int64   `json:"growth"`                  // 可以活动的成长值
	PromotionInfo         string  `json:"promotion_info"`          // 活动信息
	BillType              int64   `json:"bill_type"`               // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string  `json:"bill_header"`             // 发票抬头
	BillContent           string  `json:"bill_content"`            // 发票内容
	BillReceiverPhone     string  `json:"bill_receiver_phone"`     // 收票人电话
	BillReceiverEmail     string  `json:"bill_receiver_email"`     // 收票人邮箱
	ReceiverName          string  `json:"receiver_name"`           // 收货人姓名
	ReceiverPhone         string  `json:"receiver_phone"`          // 收货人电话
	ReceiverPostCode      string  `json:"receiver_post_code"`      // 收货人邮编
	ReceiverProvince      string  `json:"receiver_province"`       // 省份/直辖市
	ReceiverCity          string  `json:"receiver_city"`           // 城市
	ReceiverRegion        string  `json:"receiver_region"`         // 区
	ReceiverDetailAddress string  `json:"receiver_detail_address"` // 详细地址
	Note                  string  `json:"note"`                    // 订单备注
	ConfirmStatus         int64   `json:"confirm_status"`          // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int64   `json:"delete_status"`           // 删除状态：0->未删除；1->已删除
	UseIntegration        int64   `json:"use_integration"`         // 下单时使用的积分
}

type AddOrderResp struct {
}
