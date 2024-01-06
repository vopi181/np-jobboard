import { cookies } from "next/headers";
import getRequestClient from "@/app/lib/getRequestClient";

export async function POST(req: Request) {
  const data = await req.formData();
  const token = (data.get("token") as string) || "token";
  try {
    const client = getRequestClient();
    const response = await client.auth.Login({ token });
    cookies().set("auth-token", response.token);
  } catch (error) {
    console.error(error);
  }
  
  return Response.redirect(new URL(req.url).origin);
}
