
/* Rijndael.h - Sol 602963 SS8.0 gmlls 11/2011 */

#ifndef __RIJNDAEL_H__
#define __RIJNDAEL_H__

#if (!defined(SS_UNIX))
#include <exception>
#endif
#include <cstring>
#include <stdio.h>

/* Begin Infor additions for SS 8.0 11/2011 */
#include <stdlib.h>
#ifndef _ssutil_inc
#include "ssutil.inc"
#endif
/* KB 1354384 gmlls 8.0.01 12/2012 - Added additional KEY values to prevent error psc00500, which occurs 
   when the data cannot be encrypted becasue of a NULL resulting in one of the bit-shifts performed
   by the algorithm. Changing the KEY value allows additional chances to encrypt and seems to resolve 
   the issue. */
#define KEY0 "In4Gl0b@lSoLut1on&@13560Goutham@" /* Internal key used in encryption/decryption - do not change */
#define KEY1 "In4Gl0b@lSoLut1on&@13561Goutham*" /* Internal key used in encryption/decryption - do not change */
#define KEY2 "In4Gl0b@lSoLut1on&@13562Goutham&" /* Internal key used in encryption/decryption - do not change */
#define KEY3 "In4Gl0b@lSoLut1on&@13563Goutham$" /* Internal key used in encryption/decryption - do not change */
#define KEY4 "In4Gl0b@lSoLut1on&@13564Goutham!" /* Internal key used in encryption/decryption - do not change */
#define KEY5 "In4Gl0b@lSoLut1on&@13565Goutham#" /* Internal key used in encryption/decryption - do not change */
#define KEY6 "In4Gl0b@lSoLut1on&@13566Goutham%" /* Internal key used in encryption/decryption - do not change */
#define KEY7 "In4Gl0b@lSoLut1on&@13567Goutham^" /* Internal key used in encryption/decryption - do not change */
#define KEY8 "In4Gl0b@lSoLut1on&@13568Goutham(" /* Internal key used in encryption/decryption - do not change */
#define KEY9 "In4Gl0b@lSoLut1on&@13569Goutham)" /* Internal key used in encryption/decryption - do not change */
/* End KB 1354384 */
#define ENCRYPTEDLENGTH 128
/* End Infor additions for SS 8.0 11/2011 */

using namespace std;

extern "C"
{
/* Rijndael (pronounced Reindaal) is a block cipher, designed by Joan Daemen and Vincent Rijmen as a candidate algorithm for the AES.
   The cipher has a variable block length and key length. The authors currently specify how to use keys with a length
   of 128, 192, or 256 bits to encrypt blocks with al length of 128, 192 or 256 bits (all nine combinations of
   key length and block length are possible). Both block length and key length can be extended very easily to
    multiples of 32 bits.
   Rijndael can be implemented very efficiently on a wide range of processors and in hardware. 
   This implementation is based on the Java Implementation used with the Cryptix toolkit found at:
   http://www.esat.kuleuven.ac.be/~rijmen/rijndael/rijndael.zip
   Java code authors: Raif S. Naffah, Paulo S. L. M. Barreto
   This Implementation was tested against KAT test published by the authors of the method and the
   results were identical. */
class CRijndael
{
public:
	/* Operation Modes
	   The Electronic Code Book (ECB), Cipher Block Chaining (CBC) and Cipher Feedback Block (CFB) modes
	   are implemented.
	   In ECB mode if the same block is encrypted twice with the same key, the resulting
	   ciphertext blocks are the same.
	   In CBC Mode a ciphertext block is obtained by first xoring the
	   plaintext block with the previous ciphertext block, and encrypting the resulting value.
	   In CFB mode a ciphertext block is obtained by encrypting the previous ciphertext block
	   and xoring the resulting value with the plaintext. */
	enum { ECB=0, CBC=1, CFB=2 };

private:
	enum { DEFAULT_BLOCK_SIZE=16 };
	enum { MAX_BLOCK_SIZE=32, MAX_ROUNDS=14, MAX_KC=8, MAX_BC=8 };

	/* Auxiliary Functions
	   Multiply two elements of GF(2^m) */
	static int Mul(int a, int b)
	{
		return (a != 0 && b != 0) ? sm_alog[(sm_log[a & 0xFF] + sm_log[b & 0xFF]) % 255] : 0;
	}

	/* Convenience method used in generating Transposition Boxes */
	static int Mul4(int a, char b[])
	{
		if(a == 0)
			return 0;
		a = sm_log[a & 0xFF];
		int a0 = (b[0] != 0) ? sm_alog[(a + sm_log[b[0] & 0xFF]) % 255] & 0xFF : 0;
		int a1 = (b[1] != 0) ? sm_alog[(a + sm_log[b[1] & 0xFF]) % 255] & 0xFF : 0;
		int a2 = (b[2] != 0) ? sm_alog[(a + sm_log[b[2] & 0xFF]) % 255] & 0xFF : 0;
		int a3 = (b[3] != 0) ? sm_alog[(a + sm_log[b[3] & 0xFF]) % 255] & 0xFF : 0;
		return a0 << 24 | a1 << 16 | a2 << 8 | a3;
	}

public:
	/* CONSTRUCTOR */
	CRijndael();

	/* DESTRUCTOR */
	virtual ~CRijndael();

	/* Expand a user-supplied key material into a session key.
	    key        - The 128/192/256-bit user-key to use.
	    chain      - initial chain block for CBC and CFB modes.
	    keylength  - 16, 24 or 32 bytes
	    blockSize  - The block size in bytes of this Rijndael (16, 24 or 32 bytes). */
	void MakeKey(char const* key, char const* chain, int keylength=DEFAULT_BLOCK_SIZE, int blockSize=DEFAULT_BLOCK_SIZE);

private:
	/* Auxiliary Function */
	void Xor(char* buff, char const* chain)
	{
		if(false==m_bKeyInit)
#if (!defined(SS_UNIX))
			throw exception(sm_szErrorMsg1);
#else
			printf("Exception detected: %s",sm_szErrorMsg1);
#endif
		for(int i=0; i<m_blockSize; i++)
			*(buff++) ^= *(chain++);	
	}

	/* Convenience method to encrypt exactly one block of plaintext, assuming
	   Rijndael's default block size (128-bit).
	    in         - The plaintext
	    result     - The ciphertext generated from a plaintext using the key */
	void DefEncryptBlock(char const* in, char* result);

	/* Convenience method to decrypt exactly one block of plaintext, assuming
	   Rijndael's default block size (128-bit).
	    in         - The ciphertext.
	    result     - The plaintext generated from a ciphertext using the session key. */
	void DefDecryptBlock(char const* in, char* result);

public:
	/* Encrypt exactly one block of plaintext.
	    in           - The plaintext.
        result       - The ciphertext generated from a plaintext using the key. */
    void EncryptBlock(char const* in, char* result);
	
	/* Decrypt exactly one block of ciphertext.
	    in         - The ciphertext.
	    result     - The plaintext generated from a ciphertext using the session key. */
	void DecryptBlock(char const* in, char* result);

	void Encrypt(char const* in, char* result, size_t n, int iMode=ECB);
	
	void Decrypt(char const* in, char* result, size_t n, int iMode=ECB);

	/* Get Key Length */
	int GetKeyLength()
	{
		if(false==m_bKeyInit)
#if (!defined(SS_UNIX))
			throw exception(sm_szErrorMsg1);
#else
			printf("Exception detected: %s",sm_szErrorMsg1);
#endif
		return m_keylength;
	}

	/* Block Size */
	int	GetBlockSize()
	{
		if(false==m_bKeyInit)
#if (!defined(SS_UNIX))
			throw exception(sm_szErrorMsg1);
#else
			printf("Exception detected: %s",sm_szErrorMsg1);
#endif
		return m_blockSize;
	}
	
	/* Number of Rounds */
	int GetRounds()
	{
		if(false==m_bKeyInit)
#if (!defined(SS_UNIX))
			throw exception(sm_szErrorMsg1);
#else
			printf("Exception detected: %s",sm_szErrorMsg1);
#endif
		return m_iROUNDS;
	}

	void ResetChain()
	{
		memcpy(m_chain, m_chain0, m_blockSize);
	}

public:
	/* Null chain */
	static char const* sm_chain0;

private:
	static const int sm_alog[256];
	static const int sm_log[256];
	static const char sm_S[256];
    static const char sm_Si[256];
    static const int sm_T1[256];
    static const int sm_T2[256];
    static const int sm_T3[256];
    static const int sm_T4[256];
    static const int sm_T5[256];
    static const int sm_T6[256];
    static const int sm_T7[256];
    static const int sm_T8[256];
    static const int sm_U1[256];
    static const int sm_U2[256];
    static const int sm_U3[256];
    static const int sm_U4[256];
    static const char sm_rcon[30];
    static const int sm_shifts[3][4][2];
	/* Error Messages */
	static char const* sm_szErrorMsg1;
	static char const* sm_szErrorMsg2;
	/* Key Initialization Flag */
	bool m_bKeyInit;
	/* Encryption (m_Ke) round key */
	int m_Ke[MAX_ROUNDS+1][MAX_BC];
	/* Decryption (m_Kd) round key */
    int m_Kd[MAX_ROUNDS+1][MAX_BC];
	/* Key Length */
	int m_keylength;
	/* Block Size */
	int	m_blockSize;
	/* Number of Rounds */
	int m_iROUNDS;
	/* Chain Block */
	char m_chain0[MAX_BLOCK_SIZE];
	char m_chain[MAX_BLOCK_SIZE];
	/* Auxiliary private use buffers */
	int tk[MAX_KC];
	int a[MAX_BC];
	int t[MAX_BC];
};

} /* Extern C */
#endif /*  __RIJNDAEL_H__ */

