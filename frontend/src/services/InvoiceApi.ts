import BaseApi from "./BaseApi";
import InvoiceType from "@/models/InvoiceType";

class InvoiceApi extends BaseApi<InvoiceType> {
  protected resourceGroupName = "invoice";
}

const instance = new InvoiceApi();

export default instance;
