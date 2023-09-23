export interface LicensesModel {
    id:         string;
    expired_at: string;
    user_id:    string;
    created_at: string;
    features:   Features;
    type:       string;
}

export interface Features {
    rate_limit: number;
    trainings:  number;
}


export interface LicenseUsage {
    year:            number;
    license_id:      string;
    monthly_history: MonthlyHistory[];
    id:              string;
}

export interface MonthlyHistory {
    rate_limit: number;
    trainings:  number;
    month:      number;
}
