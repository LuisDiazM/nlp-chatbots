import { LicenseUsage } from "../../../shared/utilities/models";

const monthByIndex = (month: number): string => {
  if (month >= 1 && month <= 12) {
    const months = [
      "enero",
      "febrero",
      "marzo",
      "abril",
      "mayo",
      "junio",
      "julio",
      "agosto",
      "septiembre",
      "octubre",
      "noviembre",
      "diciembre",
    ];
    return months[month - 1];
  }
  return ""
};

export const createDatasetsFromLicenseUsageRateLimits = (
  data: LicenseUsage | undefined
) => {
  if (data) {
    return data?.monthly_history?.map((month) => {
      return { x: monthByIndex(month.month), y: month.rate_limit };
    });
  }
  return [];
};

export const createDatasetsFromLicenseUsageTrainigs = (
  data: LicenseUsage | undefined
) => {
  if (data) {
    return data?.monthly_history?.map((month) => {
      return { x: monthByIndex(month.month), y: month.trainings };
    });
  }
  return [];
};
