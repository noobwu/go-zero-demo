import React, {useEffect} from 'react';
import { Form, Input, Modal, Select } from 'antd';
import { LevelListItem } from '../data.d';

export interface UpdateFormProps {
  onCancel: () => void;
  onSubmit: (values: Partial<LevelListItem>) => void;
  updateModalVisible: boolean;
  currentData: Partial<LevelListItem>;
}
const FormItem = Form.Item;

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const UpdateLevelForm: React.FC<UpdateFormProps> = (props) => {
  const [form] = Form.useForm();
  const { Option } = Select;

  const {
    onSubmit,
    onCancel,
    updateModalVisible,
    currentData,
  } = props;

  useEffect(() => {
    if (form && !updateModalVisible) {
      form.resetFields();

    }
  }, [props.updateModalVisible]);

  useEffect(() => {
    if (currentData) {
      form.setFieldsValue({
        ...currentData,
      });
    }
  }, [props.currentData]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const handleFinish = (values: { [key: string]: any }) => {
    if (onSubmit) {
      onSubmit(values);
    }
  };

  const renderContent = () => {
    return (
      <>
        <FormItem
          name="id"
          label="主键"
          hidden
        >
          <Input id="update-id" placeholder="请输入主键" />
        </FormItem>
        <FormItem
          name="name"
          label="会员名"
        >
          <Input id="update-name" placeholder={'请输入会员名'}/>
        </FormItem>
        <FormItem
          name="growthPoint"
          label="成长值"
        >
          <Input id="update-growthPoint" placeholder={'请输入成长值'}/>
        </FormItem>
        <FormItem
          name="defaultStatus"
          label="是否为默认等级"
        >
          <Input id="update-defaultStatus" placeholder={'请输入是否为默认等级'}/>
        </FormItem>
        <FormItem
          name="freeFreightPoint"
          label="免运费标准"
        >
          <Input id="update-freeFreightPoint" placeholder={'请输入免运费标准'}/>
        </FormItem>
        <FormItem
          name="commentGrowthPoint"
          label="每次评价获取的成长值"
        >
          <Input id="update-commentGrowthPoint" placeholder={'请输入每次评价获取的成长值'}/>
        </FormItem>

        <FormItem
          name="status"
          label="免邮特权"
        >
          <Select id="status" placeholder={'请选择免邮特权'}>
            <Option value={0}>否</Option>
            <Option value={1}>是</Option>
          </Select>
        </FormItem>

        <FormItem
          name="priviledgeSignIn"
          label="签到特权"
        >
          <Select id="priviledgeSignIn" placeholder={'请选择签到特权'}>
            <Option value={0}>否</Option>
            <Option value={1}>是</Option>
          </Select>
        </FormItem>

        <FormItem
          name="priviledgeComment"
          label="评论获奖励特权"
        >
          <Select id="priviledgeComment" placeholder={'请选择评论获奖励特权'}>
            <Option value={0}>否</Option>
            <Option value={1}>是</Option>
          </Select>
        </FormItem>

        <FormItem
          name="priviledgePromotion"
          label="专享活动特权"
        >
          <Select id="priviledgePromotion" placeholder={'请选择专享活动特权'}>
            <Option value={0}>否</Option>
            <Option value={1}>是</Option>
          </Select>
        </FormItem>

        <FormItem
          name="priviledgeMemberPrice"
          label="会员价格特权"
        >
          <Select id="priviledgeMemberPrice" placeholder={'请选择会员价格特权'}>
            <Option value={0}>否</Option>
            <Option value={1}>是</Option>
          </Select>
        </FormItem>
      </>
    );
  };


  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel };

  return (
    <Modal
      forceRender
      destroyOnClose
      title="修改会员等级"
      visible={updateModalVisible}
      {...modalFooter}
    >
      <Form
        {...formLayout}
        form={form}
        onFinish={handleFinish}
      >
        {renderContent()}
      </Form>
    </Modal>
  );
};

export default UpdateLevelForm;
