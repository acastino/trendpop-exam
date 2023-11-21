const apiBaseUri = import.meta.env.VITE_BACKEND_API_BASEURL;

const headers = {
  Accept: "application/json",
  "Content-Type": "application/json",
};

export default abstract class BaseApi<ModelType> {
  protected abstract resourceGroupName: string;

  private resourceGroupUri() {
    return `${apiBaseUri}/${this.resourceGroupName}`;
  }

  private request(
    method: string,
    endpoint: string,
    data?: ModelType
  ): Promise<Response> {
    return fetch(`${this.resourceGroupUri()}${endpoint}`, {
      body: data && JSON.stringify(data),
      headers,
      method,
    });
  }

  /**
   * create
   */
  public create(data: ModelType): Promise<Response> {
    return this.request("post", "", data);
  }
}
