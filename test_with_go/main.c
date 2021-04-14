#include<stdio.h>
#include "ecc.c"

int main() {
    uint8_t pub_key[secp256r1+1] = {
        0x02, 0x68, 0xc0, 0xc8, 0x1d, 0x72, 0x85, 0x67, 
        0x22, 0xe0, 0x37, 0x38, 0xa7, 0xb4, 0x6c, 0x11, 
        0x62, 0x85, 0xc1, 0xa3, 0xa8, 0x50, 0xee, 0xfc, 
        0x84, 0xa6, 0xe7, 0x47, 0x78, 0x1f, 0x22, 0x1d,
        0x0a
    };
    int i;
    printf("pub key:\n");
    for (i = 0; i < secp256r1+1; i++) {
        printf("%02x", pub_key[i]);
    }
    printf("\n");

    FILE *f;
    uint8_t sig[64];
    f = fopen("sig", "r");
    fread(sig, 1, 64, f);
    fclose(f);

    printf("signature:\n");
    for (i = 0; i < 64; i++) {
        printf("%02x", sig[i]);
    }
    printf("\n");

    uint8_t hash[32];
    for (i = 0; i < 32; i++) {
        hash[i] = '1';
    }

    printf("hash:\n");
    for (i = 0; i < 32; i++) {
        printf("%02x", hash[i]);
    }
    printf("\n");

    int result = ecdsa_verify(pub_key, hash, sig);
    if (result == 1) {
        printf("Valid\n");
    } else {
        printf("Invalid\n");
    }
    return 0;
}
