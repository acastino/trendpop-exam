const apis = {
  laravel: import.meta.env.VITE_LARAVEL_BACKEND_API_BASEURL,
  golang: import.meta.env.VITE_GOLANG_BACKEND_API_BASEURL,
};

const headers = {
  Accept: "application/json",
  "Content-Type": "application/json",
};

export default abstract class BaseApi<ModelType> {
  protected abstract resourceGroupName: string;

  private resourceGroupUri() {
    const queryString = location.search.substring(1);
    const params = new URLSearchParams(queryString);
    const backend = params.get("backend")?.toLowerCase() || "";
    const apiBaseUri = backend == "golang" ? apis.golang : apis.laravel;
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
