To configure the **Cloudflare API Token** for cert-manager to perform **DNS-01 challenges**, you can create a custom API token in Cloudflare with the necessary permissions. Cloudflare provides some predefined templates, but for cert-manager’s DNS-01 challenge, you'll need to create a **custom API token** with specific permissions for your DNS zone.

### Steps to Create the Cloudflare API Token:

1. **Go to Cloudflare Dashboard**:
   - Log in to your Cloudflare account.
   - Navigate to **My Profile** > **API Tokens**.

2. **Create a Custom API Token**:
   - Click on **Create Token**.
   - Scroll down to the **Custom Token** section and click on **Create Custom Token**.

3. **Configure API Token Permissions**:
   - Set the following permissions for cert-manager's DNS-01 challenge solver:
     - **Zone**: **DNS** (Edit) – This allows cert-manager to create and delete DNS TXT records required for DNS-01 challenges.
     - **Zone**: **Zone** (Read) – This gives cert-manager access to read zone details like DNS settings for your domain.

4. **Zone Resources**:
   - Choose the specific zone (your domain) that the token will have access to. For example, if your domain is `kowlon.my.id`, select **Specific zone** and then specify `kowlon.my.id`.

5. **Token Expiration (Optional)**:
   - You can set an expiration for the token if needed, but it is generally not necessary unless you want to enforce time-limited access.

6. **Create Token**:
   - After configuring the permissions, click on **Continue to Summary**, review the permissions, and then click **Create Token**.
   - **Save the token securely**. You'll use this token in your Kubernetes secret for cert-manager.

### Permissions Overview for cert-manager:
| Resource | Permission | Description |
|----------|-------------|-------------|
| **Zone** | **DNS** (Edit) | Allows cert-manager to create/delete TXT records for DNS-01 challenges. |
| **Zone** | **Zone** (Read) | Allows cert-manager to read DNS zone details to manage the records. |

### Example Configuration:

Once you've created the token, store it in a Kubernetes secret for cert-manager to use:

```bash
kubectl create secret generic cloudflare-api-token-secret \
  --from-literal=api-token=<your-cloudflare-api-token> \
  --namespace <namespace>
```

### Usage in `ClusterIssuer`:

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: production-cluster-issuer
spec:
  acme:
    email: your-email@example.com
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
      - dns01:
          cloudflare:
            apiTokenSecretRef:
              name: cloudflare-api-token-secret
              key: api-token
```

By following these steps, cert-manager will have the required permissions to create DNS-01 challenge records in Cloudflare, allowing for the successful issuance of certificates.