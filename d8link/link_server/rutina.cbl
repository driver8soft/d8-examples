       IDENTIFICATION DIVISION.
       PROGRAM-ID. rutina.
       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       DATA DIVISION.
       FILE SECTION.
       WORKING-STORAGE SECTION.
      * In COBOL, you declare variables in the WORKING-STORAGE section

       LINKAGE SECTION.
      * Data to share with COBOL subroutines 
       01 TOTAL PIC S9(9).
       PROCEDURE DIVISION USING TOTAL. 
      * code goes here!
           ADD 1 TO TOTAL.
           GOBACK.



