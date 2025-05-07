import request from "@/utils/request";
// API基础URL
const API_BASE_URL = "/api/v1/apis";

const ApiAPI = {
  /**
   * 获取API下拉数据源
   * @returns {Promise} API下拉数据源
   */
  getOptions() {
    return request({
      url: `${API_BASE_URL}/options`,
      method: "get",
    });
  },
};

export default ApiAPI;
