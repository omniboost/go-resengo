package fortnox

type Vouchers []Voucher

type Voucher struct {
	URL             string        `json:"@url,omitempty"`            // Direct URL to the record.
	Comments        string        `json:"Comments"`                  // Comments of the voucher.
	CostCenter      string        `json:"CostCenter"`                // Code of the cost center. The code must be of an existing cost center.
	Description     string        `json:"Description"`               // Description of the voucher.
	Project         string        `json:"Project"`                   // Code of the project. The code must be of an existing project.
	ReferenceNumber string        `json:"ReferenceNumber,omitempty"` // Reference number, for example an invoice number.
	ReferenceType   string        `json:"ReferenceType"`             // Reference type. Can be INVOICE SUPPLIERINVOICE INVOICEPAYMENT SUPPLIERPAYMENT MANUAL CASHINVOICE or ACCRUAL
	TransactionDate Date          `json:"TransactionDate"`           // Date of the transaction. Must be a valid date according to our date format.
	VoucherNumber   int           `json:"VoucherNumber,omitempty"`   // Number of the voucher
	VoucherRows     VoucherRows   `json:"VoucherRows"`               // The properties for the object in this array is listed in the table for “Voucher Rows”.
	VoucherSeries   string        `json:"VoucherSeries"`             // Code of the voucher series. The code must be of an existing voucher series.
	Year            int           `json:"Year,omitempty"`            // Id of the year of the voucher.
	ApprovalState   ApprovalState `json:"ApprovalState,omitempty"`   // // The approval state f the voucher.
}

type VoucherRows []VoucherRow

type VoucherRow struct {
	Account                int     `json:"Account"`                // Account number. The number must be of an existing active account.
	CostCenter             string  `json:"CostCenter"`             // Code of the cost center. The code must be of an existing cost center.
	Credit                 float64 `json:"Credit"`                 // Amount of credit.
	Description            string  `json:"Description"`            // The description of the account.
	Debit                  float64 `json:"Debit"`                  // Amount of debit.
	Project                string  `json:"Project"`                // Code of the project. The code must be of an existing project.
	Removed                bool    `json:"Removed,omitempty"`      // If the row is marked as removed.
	TransactionInformation string  `json:"TransactionInformation"` // Transaction information regarding the row.
	Quantity               int     `json:"Quantity,omitempty"`
}

type ApprovalState int

var (
	ApprovalStateNotForApproval     ApprovalState = 0
	ApprovalStateNotReadForApproval ApprovalState = 1
	ApprovalStateNotApproved        ApprovalState = 2
	ApprovalStateApproved           ApprovalState = 3
)
