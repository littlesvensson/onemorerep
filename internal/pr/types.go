package pr

type EstimateRequest struct {
    Weight float64 `json:"weight" binding:"required,gt=0"`
    Reps   int     `json:"reps"   binding:"required,min=1,max=30"`
}

type EstimateResponse struct {
    Input   EstimateRequest    `json:"input"`
    OneRM   float64            `json:"one_rm"`
    Table   map[string]float64 `json:"rep_table"`
    Formula string             `json:"formula"`
    Notes   []string           `json:"notes,omitempty"`
}

type ErrResp struct {  
    Error string `json:"error"`
}
